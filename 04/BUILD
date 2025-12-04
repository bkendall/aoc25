load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "04_lib",
    srcs = ["main.go"],
    importpath = "aoc25/04",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "04",
    embed = [":04_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "04_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":04_lib"],
)
