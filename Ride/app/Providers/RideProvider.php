<?php

namespace App\Providers;
use app\Service\Ride;
use Illuminate\Support\ServiceProvider;

class RideProvider extends ServiceProvider
{
    /**
     * Register services.
     */
    public function register(): void
    {
         $this->app->singleton(Ride::class, function($app){
            return new Ride(
                $app->make(\App\Service\User::class),
                 $app->make(\App\Service\RideRepository::class),
                 $app->make(\App\Service\ExternalUserRepository::class)
            );
        });
    }

    /**
     * Bootstrap services.
     */
    public function boot(): void
    {
        //
    }
}
