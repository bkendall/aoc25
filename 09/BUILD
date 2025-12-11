load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "09_lib",
    srcs = ["main.go"],
    importpath = "aoc25/09",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "09",
    embed = [":09_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "09_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":09_lib"],
)
