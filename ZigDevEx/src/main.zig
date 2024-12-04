const std = @import("std");
const zap = @import("zap");

fn default_handler(r: zap.Request) void {
    if (r.path) |the_path| {
        std.debug.print("PATH: {s}\n", .{the_path});
    }

    if (r.query) |the_query| {
        std.debug.print("QUERY: {s}\n", .{the_query});
    }

    r.sendBody("<html><body><h1>Hello World!</h1></body></html>") catch return;
}

//run on localhost:9000
pub fn main() !void {
    var listener = zap.HttpListener.init(.{
        .port = 9000,
        .on_request = default_handler,
        .log = true,
        .max_clients = 100000,
    });
    try listener.listen();

    std.debug.print("Listening on http://127.0.0.1:9000\n", .{});

    zap.start(.{
        .threads = 2,
        .workers = 1, // 1 worker enables sharing state between threads
    });
}
