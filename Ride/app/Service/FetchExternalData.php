<?php

namespace App\Service;

class FetchExternalData
{
    /**
     * Create a new class instance.
     */
    public function __construct()
    {
        //
    }

    public function getInfoBasicAuth($url, $username, $password)
    {
        $client = new \GuzzleHttp\Client();
        try {
            $response = $client->request('GET', $url, [
                'auth' => [$username, $password]
            ]);

            return json_decode($response->getBody(), true);
        } catch (\Exception $e) {
            // Handle exceptions (e.g., log the error)
            return null;
        }
    }
}

//  protected UserValidation $userValidation;

//     protected GetOnlineDriver $getOnlineDriver;

//     public function __construct(
//         ?UserValidation $userValidation = null,
//         ?GetOnlineDriver $getOnlineDriver = null
//     ) {
//         $baseUrl = config('services.external_api.url', 'http://localhost:8080');
//         $username = config('services.external_api.username');
//         $password = config('services.external_api.password');

//         $this->userValidation = $userValidation ?? new UserValidation($baseUrl, $username, $password);
//         $this->getOnlineDriver = $getOnlineDriver ?? new GetOnlineDriver($baseUrl, $username, $password);
//     }

//     public function validateUser(int $userId, string $type = 'rider'): array
//     {
//         return $this->userValidation->validateUser($userId, $type);
//     }
//     public function getOnlineDrivers(): ?array
//     {
//         return $this->getOnlineDriver->getOnlineDrivers();
//     }
