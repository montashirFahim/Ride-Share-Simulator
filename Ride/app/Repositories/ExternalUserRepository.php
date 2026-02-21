<?php

namespace App\Repositories;
use Illuminate\Support\Facades\Http;

class ExternalUserRepository
{
    public function __construct()
    {
    }

    public function getWithBasicAuth($url, $username, $password)
    {
        try {
            return Http::withBasicAuth($username, $password)->timeout(3)->get($url);
        } catch (\Exception $e) {
            return null;
        }
    }
}
