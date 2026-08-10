if "%~1" == "--start" (
    if "%~2" == "-b" (
        docker compose --env-file ./config.env up --build
    ) else (
        docker compose --env-file ./config.env up
    )
) else if "%~1" == "--stop" (
    docker compose down
) else if "%~1" == "--restart" (
    docker compose down

    if "%~2" == "-b" (
        docker compose --env-file ./config.env up --build
    ) else (
        docker compose --env-file ./config.env up
    )
) 