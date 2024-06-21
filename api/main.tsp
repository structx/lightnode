
import "@typespec/http";

using TypeSpec.Http;

model PartialBlock {}

model Block {}

model PartialTx {}

model Tx {}

@route("/api)
namespace api {

    @route("/v1)
    namespace v1 {

        @route("/blocks") 
        namespace blocks {
            op read(@path hash: string): Block;
            op list()
        }
    }
}