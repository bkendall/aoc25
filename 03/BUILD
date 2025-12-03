load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "03_lib",
    srcs = ["main.go"],
    importpath = "aoc25/03",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "03",
    embed = [":03_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "03_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":03_lib"],
)
