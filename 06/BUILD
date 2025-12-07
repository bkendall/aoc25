load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "06_lib",
    srcs = ["main.go"],
    importpath = "aoc25/06",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "06",
    embed = [":06_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "06_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":06_lib"],
)
