<?php

namespace App\Services;

use App\Repositories\ExternalUserRepository;
use App\Repositories\RideRepository;

class RideService
{

    public function __construct(
        protected UserService $userService,
        protected RideRepository $rideRepo,
        protected ExternalUserRepository $externalUserRepo
    ) {
    }
    public function validateUser($userId, $type)
    {
        $response = $this->userService->getUserInfo($userId);
        if (!$response)
            return 2;
        else if ($response->json('user_type') != $type)
            return 0;
        return 1;
    }

    public function checkActiveRide($userId)
    {
        $rideId = $this->rideRepo->FindUserActiveRide($userId);

        return $rideId ? true : false;
    }

    public function assignDriver()
    {
        $response = $this->userService->getOnlineDriver();

        if (!$response)
            return -1;

        $drivers = json_decode($response, true);

        if (!$drivers || !is_array($drivers))
            return -1;

        $busyDrivers = $this->rideRepo->GetBusyDrivers();

        $busyMap = [];
        foreach ($busyDrivers as $busy) {
            $busyMap[$busy->driver_id] = true;
        }

        foreach ($drivers as $driver) {
            if (!isset($busyMap[$driver['id']])) {
                return $driver['id'];
            }
        }

        return -1;
    }

    public function newRideCreate($riderId, $driverId)
    {
        return $this->rideRepo->CreateRide($riderId, $driverId);
    }

    public function checkValidRideById($rideId)
    {
        $res = $this->rideRepo->FindRideId($rideId);

        if (!$res)
            return false;
        return true;
    }

    public function endRide($rideId)
    {
        return $this->rideRepo->EndRide($rideId);
    }

    public function getRideInfo($userId)
    {
        return $this->rideRepo->GetRideInfo($userId);
    }

    public function checkEndRide($rideId)
    {
        $res = $this->rideRepo->EndRideCheck($rideId);
        if (!$res)
            return false;
        return true;
    }
}