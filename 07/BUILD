load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "07_lib",
    srcs = ["main.go"],
    importpath = "aoc25/07",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "07",
    embed = [":07_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "07_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":07_lib"],
)
