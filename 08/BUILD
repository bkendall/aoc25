load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "08_lib",
    srcs = ["main.go"],
    importpath = "aoc25/08",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "08",
    embed = [":08_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "08_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":08_lib"],
)
