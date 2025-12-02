load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "01_lib",
    srcs = ["main.go"],
    importpath = "aoc25/01",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "01",
    embed = [":01_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)
