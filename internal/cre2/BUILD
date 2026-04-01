# Bazel build for cre2 C/C++ wrapper
# Exports source files and headers for downstream consumers
# This is a filegroup, not a compiled library - compilation happens in consuming repo

filegroup(
    name = "cre2_wrapper",
    srcs = [
        "cre2.cpp",
        "cre2.h",
    ],
    visibility = ["//visibility:public"],
)
