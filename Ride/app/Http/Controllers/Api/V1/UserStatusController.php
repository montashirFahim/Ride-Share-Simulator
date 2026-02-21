<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Services\RideService;
use App\Services\UserService;

class UserStatusController extends Controller
{
    protected RideService $rideService;
    protected UserService $userService;

    public function __construct(RideService $rideService, UserService $userService)
    {
        $this->rideService = $rideService;
        $this->userService = $userService;
    }

    public function getUserStatus(Request $request)
    {
        $riderId = $request->query('rider_id');
        $driverId = $request->query('driver_id');

        if ($riderId)
            return $this->getRiderStatus($riderId);
        return $this->getDriverStatus($driverId);
    }

    protected function getRiderStatus($riderId)
    {
        // Check Valid Rider
        $userInfo = $this->userService->getUserInfo($riderId);
        if (!$userInfo)
            return response()->json(['message' => 'Rider Not Found'], 404);
        $userInfo = json_decode($userInfo, true);

        //Check Active Ride 
        $rideStatus = $this->rideService->checkActiveRide($riderId);
        if (!$rideStatus)
            return response()->json(['message' => 'No active ride found'], 200);

        // Get rider info
        $rideInfo = $this->rideService->getRideInfo($riderId);
        if (!$rideInfo)
            return response()->json(['message' => 'Error fetching ride info'], 404);

        //return response
        return response()->json([
            "id" => (string) $rideInfo->id,
            "driver_id" => (string) $rideInfo->driver_id,
            "rider_id" => (string) $riderId,
            "rider_phone" => (string) $userInfo['mobile_no'],
            "status" => "started"
        ], 200);
    }

    public function getDriverStatus($driverId)
    {
        // Check Valid Driver
        $userInfo = $this->userService->getUserInfo($driverId);
        if (!$userInfo)
            return response()->json(['message' => 'Driver Not Found'], 404);

        $userInfo = json_decode($userInfo, true);
        // // Check Active Ride 
        $rideStatus = $this->rideService->checkActiveRide($driverId);
        if (!$rideStatus)
            return response()->json(['message' => 'No active ride found'], 200);

        $rideInfo = $this->rideService->getRideInfo($driverId);
        if (!$rideInfo)
            return response()->json(['message' => 'Error fetching ride info'], 404);
        //return response
        return response()->json([
            "id" => (string) $rideInfo->id,
            "driver_id" => (string) $driverId,
            "rider_id" => (string) $rideInfo->rider_id,
            "driver_phone" => (string) $userInfo['mobile_no'],
            "status" => "started"
        ], 200);
    }
}
