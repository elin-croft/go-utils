namespace py ew_toolbox
namespace go ew_go_thrift.ew_toolbox

struct request{
    1: i64 id;
}
struct response{
    1: i64 id;
    2: string name;
}
service predict{
    response get(1:request req);
}
