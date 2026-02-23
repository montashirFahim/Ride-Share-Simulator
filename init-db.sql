-- Create userdb and riderdb databases
CREATE DATABASE userdb;
CREATE DATABASE riderdb;

\c userdb

-- Create enum types for users table
CREATE TYPE user_type_enum AS ENUM ('rider', 'driver');
CREATE TYPE user_status_enum AS ENUM ('online', 'offline');

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    mobile_no VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    user_type user_type_enum NOT NULL,
    cur_status user_status_enum DEFAULT 'offline' NOT NULL,
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for users table
CREATE INDEX IF NOT EXISTS idx_users_mobile ON users(mobile_no);



\c riderdb

-- Create enum type for ride status
CREATE TYPE status_type AS ENUM ('started', 'ended');

-- Create rides table
CREATE TABLE IF NOT EXISTS rides (
    id SERIAL PRIMARY KEY,
    rider_id INTEGER NOT NULL,
    driver_id INTEGER NOT NULL,
    status status_type DEFAULT 'started' NOT NULL,
    started_at TIMESTAMP NOT NULL,
    ended_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
