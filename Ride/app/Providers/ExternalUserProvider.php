<?php

namespace App\Providers;
use app\Service\ExternalUserRepository;
use Illuminate\Support\ServiceProvider;


class ExternalUserProvider extends ServiceProvider
{
    /**
     * Register services.
     */
    public function register(): void
    {
        $this->app->singleton(ExternalUserRepository::class, function($app){
            return new ExternalUserRepository();
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
