# Test task

If you have the `make` utility on your PC, you can just execute `FILE=<data_file> make` to build the **Docker** image and run the application with the specified data file.

Otherwise, run the following commands:

`docker build -t task:0.0.1 .`

`docker run --mount type=bind,source=$(pwd)/<data_file>,target=/bin/<data_file> task:0.0.1 /bin/<data_file>`

The data file should be a **JSON** or **CSV** file.