load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "10_lib",
    srcs = ["main.go"],
    importpath = "aoc25/10",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "10",
    embed = [":10_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "10_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":10_lib"],
    deps = ["@com_github_google_go_cmp//cmp"],
)
