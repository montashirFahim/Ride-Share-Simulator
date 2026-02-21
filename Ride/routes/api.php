<?php

use App\Http\Controllers\Api\V1\ApiController;
use Illuminate\Support\Facades\Route;

Route::prefix("v1")->group(function(){
    Route::post("rides", [ApiController::class, "getRide"]);
    Route::put("rides",[ApiController::class, "endRide"]);
    Route::get("rides/{rider_id?}",[ApiController::class,"getRiderStatus"]);
    Route::get("rides/{driver_id?}",[ApiController::class,"getDriverStatus"]);
});