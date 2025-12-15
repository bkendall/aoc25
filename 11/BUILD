load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "11_lib",
    srcs = ["main.go"],
    importpath = "aoc25/11",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "11",
    embed = [":11_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "11_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":11_lib"],
    deps = ["@com_github_google_go_cmp//cmp"],
)
