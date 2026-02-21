<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Services\RideService;
use App\Services\UserService;

class RideController extends Controller
{
    protected RideService $rideService;

    public function __construct(RideService $rideService)
    {
        $this->rideService = $rideService;
    }

    public function getRide(Request $request)
    {
        if (empty(request()->get("rider_id")))
            return response()->json(['message' => 'Rider ID is required'], 400);

        $userId = (int) request()->get("rider_id");

        //user validation 
        $userValidation = $this->rideService->validateUser($userId, "rider");
        if ($userValidation == 0)
            return response()->json(['message' => 'Invalid id or user not found'], 404);
        else if ($userValidation == 2)
            return response()->json(['message' => 'Internal Server Error'], 500);

        //check if user is on an active ride
        $activeRide = $this->rideService->checkActiveRide($userId);
        if ($activeRide)
            return response()->json(['message' => 'User has alreay an active ride'], 403);

        //Assign Driver
        $driverId = $this->rideService->assignDriver();
        if ($driverId == -1)
            return response()->json(['message' => 'No active driver found'], 404);

        //create a new ride
        $rideId = $this->rideService->newRideCreate($userId, $driverId);
        if (!$rideId)
            return response()->json(['message' => 'Error creating ride'], 500);

        return response()->json(['id' => (string) $rideId, 'rider_id' => (string) $userId, 'driver_id' => (string) $driverId, 'status' => 'started'], 200);
    }

    public function endRide(Request $request)
    {
        if ($this->rideService->checkEndRide($request->id))
            return response()->json(['message' => 'Ride already ended'], 403);

        //Check if ride exists and is active
        $isValidRide = $this->rideService->checkValidRideById($request->id);
        if (!$isValidRide)
            return response()->json(['message' => 'Invalid Ride ID'], 404);

        //End Ride
        $endRide = $this->rideService->endRide($request->id);
        if (!$endRide)
            return response()->json(['message' => 'Error ending ride'], 500);
        return response()->json(['id' => (string) $endRide->id, 'rider_id' => (string) $endRide->rider_id, 'driver_id' => (string) $endRide->driver_id, 'status' => (string) $endRide->status], 200);
    }
}
