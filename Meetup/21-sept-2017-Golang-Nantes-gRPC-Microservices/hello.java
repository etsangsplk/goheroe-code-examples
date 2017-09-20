@Override
protected String doInBackground(Void... nothing) {
    try {
        mChannel = ManagedChannelBuilder.forAddress(mHost, mPort)
            .usePlaintext(true)
            .build();
        GreeterGrpc.GreeterBlockingStub stub = GreeterGrpc.newBlockingStub(mChannel);
        HelloRequest message = HelloRequest.newBuilder().setName(mMessage).build();
        HelloReply reply = stub.sayHello(message);
        return reply.getMessage();
    } catch (Exception e) {
        StringWriter sw = new StringWriter();
        PrintWriter pw = new PrintWriter(sw);
        e.printStackTrace(pw);
        pw.flush();
        return String.format("Failed... : %n%s", sw);
    }
}