load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "05_lib",
    srcs = ["main.go"],
    importpath = "aoc25/05",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "05",
    embed = [":05_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "05_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":05_lib"],
)
