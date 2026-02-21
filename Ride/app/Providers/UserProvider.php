<?php

namespace App\Providers;
use app\Service\User;
use Illuminate\Support\ServiceProvider;

class UserProvider extends ServiceProvider
{
    /**
     * Register services.
     */
    public function register(): void
    {
        $this->app->singleton(User::class, function($app){
            return new User(
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
