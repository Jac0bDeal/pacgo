# Pacgo

Inspired by https://github.com/danicat/pacgo.

## Build
This project is built on Go 1.15, so Go 1.15+ is recommended.
Other than Go, the only other dependency required is Make.

To do a clean build of the project, run
```shell script
make
```

The `pacgo` binary will be built at `bin/pacgo`.

## Running
From the root of the project directory run
```shell script
./bin/pacgo
```

This will start up the game with the default config and level. To specify 
a different config, pass the desire config filepath in via the `--config-file` 
arg. To play a different level, pass in the filepath to the level file via 
the `--level-file` flag. For consistency, configs should go under the `configs/` 
directory and level files under the `levels/` directory.

## Customization
### Config

### Levels
