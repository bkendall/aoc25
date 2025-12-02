load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "02_lib",
    srcs = ["main.go"],
    importpath = "aoc25/02",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "02",
    embed = [":02_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)
