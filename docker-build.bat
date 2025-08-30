@echo off
REM ASCII Art Web Docker Build Script for Windows
REM This script builds the Docker image and provides various operations

setlocal enabledelayedexpansion

set IMAGE_NAME=ascii-art-web
set CONTAINER_NAME=ascii-art-web-container
set TAG=latest

if "%1"=="" goto usage
if "%1"=="build" goto build
if "%1"=="run" goto run
if "%1"=="stop" goto stop
if "%1"=="clean" goto clean
if "%1"=="logs" goto logs
if "%1"=="shell" goto shell
goto usage

:usage
echo Usage: %0 [build^|run^|stop^|clean^|logs^|shell]
echo.
echo Commands:
echo   build  - Build the Docker image
echo   run    - Run the container
echo   stop   - Stop the container
echo   clean  - Remove container and image
echo   logs   - Show container logs
echo   shell  - Open shell in running container
goto end

:build
echo Building Docker image...
docker build -t %IMAGE_NAME%:%TAG% .
if errorlevel 1 (
    echo Failed to build image
    goto end
)
echo Image built successfully!
echo.
echo Image details:
docker images %IMAGE_NAME%:%TAG%
goto end

:run
echo Running container...

REM Stop existing container if running
docker ps -q -f name=%CONTAINER_NAME% >nul 2>&1
if not errorlevel 1 (
    echo Stopping existing container...
    docker stop %CONTAINER_NAME%
    docker rm %CONTAINER_NAME%
)

REM Run new container
docker run -d --name %CONTAINER_NAME% -p 8080:8080 --restart unless-stopped %IMAGE_NAME%:%TAG%
if errorlevel 1 (
    echo Failed to start container
    goto end
)

echo Container started successfully!
echo Access the application at: http://localhost:8080
echo.
echo Container status:
docker ps -f name=%CONTAINER_NAME%
goto end

:stop
echo Stopping container...
docker stop %CONTAINER_NAME% 2>nul
docker rm %CONTAINER_NAME% 2>nul
echo Container stopped and removed
goto end

:clean
echo Cleaning up...
docker stop %CONTAINER_NAME% 2>nul
docker rm %CONTAINER_NAME% 2>nul
docker rmi %IMAGE_NAME%:%TAG% 2>nul
docker system prune -f
echo Cleanup completed
goto end

:logs
echo Container logs:
docker logs -f %CONTAINER_NAME%
goto end

:shell
echo Opening shell in container...
docker exec -it %CONTAINER_NAME% /bin/sh
goto end

:end
endlocal