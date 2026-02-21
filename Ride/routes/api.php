<?php

use App\Http\Controllers\Api\V1\RideController;
use App\Http\Controllers\Api\V1\UserStatusController;
use Illuminate\Support\Facades\Route;

Route::prefix("v1")->group(function () {
    Route::post("rides", [RideController::class, "getRide"]);
    Route::put("rides", [RideController::class, "endRide"]);
    Route::get("rides", [UserStatusController::class, "getUserStatus"]);
});