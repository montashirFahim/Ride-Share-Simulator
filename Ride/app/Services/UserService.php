<?php

namespace App\Services;

use App\Repositories\ExternalUserRepository;
use App\Repositories\RideRepository;

class UserService
{
    private $baseUrl, $username, $password;

    public function __construct(
        protected RideRepository $rideRepo,
        protected ExternalUserRepository $externalUserRepo
    ) {
        $this->baseUrl = config('services.external_api.url', 'http://localhost:8080');
        $this->username = config('services.external_api.username');
        $this->password = config('services.external_api.password');
    }

    public function getUserInfo($userId)
    {
        $url = rtrim($this->baseUrl, '/') . "/api/v1/users/" . $userId . "/info";

        $response = $this->externalUserRepo->getWithBasicAuth(
            $url,
            $this->username,
            $this->password
        );

        if (!$response)
            return null;

        return $response;
    }

    public function getOnlineDriver()
    {
        $url = rtrim($this->baseUrl, '/') . "/api/v1/drivers/online";

        $response = $this->externalUserRepo->getWithBasicAuth(
            $url,
            $this->username,
            $this->password
        );

        if (!$response || $response->notFound() || $response->badRequest())
            return null;

        return $response;
    }

}
