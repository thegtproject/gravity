@echo off
REM Gravity Development Toolkit Wrapper for Windows

if "%1"=="" goto Usage

go run ./developing/tools %1

goto Exit

:Usage
echo.
echo Gravity Development Toolkit Wrapper for Windows
echo.
echo usage: %0 ^<command^>
echo        %0 genversion
echo.
:Exit
