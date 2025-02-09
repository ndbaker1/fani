import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { FaniService } from '../gen/fani/v1/fani_pb'

const baseUrl = "http://localhost:8080"

export const faniClient = createClient(
    FaniService,
    createConnectTransport({ baseUrl }),
);
