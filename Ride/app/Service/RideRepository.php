<?php

namespace App\Service;

use Illuminate\Support\Facades\DB;

class RideRepository
{
    public function __construct()
    {
        //
    }

    public function CreateRide(int $riderId, int $driverId)
    {
        $result = DB::select(
            "INSERT INTO rides (rider_id, driver_id, status, started_at, created_at, updated_at)
             VALUES (?, ?, 'started', NOW(), NOW(), NOW()) RETURNING id",
            [$riderId, $driverId]
        );

        return $result[0]->id ?? null;
    }

    public function EndRide(int $rideId)
    {
        $result = DB::select(
            "UPDATE rides
             SET status = 'ended',
                 ended_at = NOW(),
                 updated_at = NOW()
             WHERE id = ?
               AND status = 'started'
             RETURNING id",
            [$rideId]
        );

        return $result[0]->id ?? null;
    }

    public function FindUserActiveRide(int $userId)
    {
        return DB::select(
            "SELECT id
             FROM rides
             WHERE (rider_id = ? OR driver_id = ?)
               AND status = 'started'",
            [$userId, $userId]
        );
    }

    public function FindRideId(int $id)
    {
        return DB::selectOne(
            "SELECT *
             FROM rides
             WHERE id = ?
             LIMIT 1",
            [$id]
        );
    }

    public function GetBusyDrivers()
    {
        return DB::select(
            "SELECT DISTINCT driver_id
             FROM rides
             WHERE status = 'started'"
        );
    }
}
