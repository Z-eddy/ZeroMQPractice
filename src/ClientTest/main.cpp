#include <iostream>
#include <string>

#include "zmq.hpp"
#include "zmq_addon.hpp"

int main() {
  // initialize the zmq context with a single IO thread
  zmq::context_t context{1};

  // construct a REQ (request) socket and connect to interface
  zmq::socket_t socket{context, zmq::socket_type::sub};
  socket.connect("tcp://192.168.0.13:9092");
  socket.set(zmq::sockopt::subscribe, "");

  while (true) {
    // Receive all parts of the message
    std::vector<zmq::message_t> recv_msgs;
    zmq::recv_result_t result =
        zmq::recv_multipart(socket, std::back_inserter(recv_msgs));

    for (const auto &item : recv_msgs) {
      std::cout << item.to_string() << std::endl;
    }
  }

  return 0;
}
