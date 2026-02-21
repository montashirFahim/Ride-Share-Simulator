<?php

namespace App\Providers;
use app\Service\RideRepository;
use Illuminate\Support\ServiceProvider;

class RideRepositoryProvider extends ServiceProvider
{
    /**
     * Register services.
     */
    public function register(): void
    {
        $this->app->singleton(RideRepository::class, function($app){
            return new RideRepository();
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
