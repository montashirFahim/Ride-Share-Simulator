<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class ApiController extends Controller
{
    public function getRide(Request $request)
    {
       // User validation
       // Check if rider is already in a ride
       // Assign driver
       //
        return response()->json(['message' => 'Tonoy Gay'], 200);
    }

    public function endRide(Request $request)
    {
        return response()->json(['message' => 'Tonoy Gay'], 200);
    }

    public function getRiderStatus(Request $request)
    {
        return response()->json(['message' => 'Tonoy Gay'], 200);
    }

    public function getDriverStatus(Request $request)
    {
        return response()->json(['message' => 'Tonoy Gay'], 200);
    }
}
