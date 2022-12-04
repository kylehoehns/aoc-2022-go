# Advent of Code 2022

Working through problems for [Advent of Code 2022](https://adventofcode.com/2022).

Using [Go](https://go.dev/)

## Running

Use VS Code and open this project in a [Dev Container](https://code.visualstudio.com/docs/devcontainers/containers) without needing to install any tooling locally.

Inside the container, run the following command to test a given day's problem.

```shell
go test ./puzzles/day00
```

To run the given puzzle input for the day.

```shell
go run puzzles/day00/main.go
```
```
          .     .  .      +     .      .          .
     .       .      .     #       .           .
        .      .         ###            .      .      .
      .      .   "#:. .:##"##:. .:#"  .      .
          .      . "####"###"####"  .
       .     "#:.    .:#"###"#:.    .:#"  .        .       .
  .             "#########"#########"        .        .
        .    "#:.  "####"###"####"  .:#"   .       .
     .     .  "#######""##"##""#######"                  .
                ."##"#####"#####"##"           .      .
    .   "#:. ...  .:##"###"###"##:.  ... .:#"     .
      .     "#######"##"#####"##"#######"      .     .
    .    .     "#####""#######""#####"    .      .
            .     "      000      "    .     .
       .         .   .   000     .        .       .
........................O000O................................
```
