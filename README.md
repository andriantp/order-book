# Order Book in Go

A step-by-step journey to building a modern exchange engine in Go.

This repository accompanies a Medium series where we gradually build an Order Book and Matching Engine from scratch, then evolve it into a distributed trading system using techniques commonly found in real-world exchanges.

Each chapter introduces one concept at a time, keeping the implementation small, focused, and easy to understand.

---

# 📚 Related Articles

## Phase 1 : Foundations

- What Is an Order Book? : [Chapter 1](https://medium.com/@andriantriputra/orderbook-1-what-is-an-order-book-dc8221b80992)
- Price-Time Priority : [Chapter 2](https://medium.com/@andriantriputra/orderbook-2-price-time-priority-86db43dfd7d3)
- Order Book Data Structures : [Chapter 3](https://medium.com/@andriantriputra/orderbook-3-order-book-data-structures-f7966db86ed8)

## Phase 2 : Core Engine

- Building the Order Book : [Chapter 4](https://medium.com/@andriantriputra/orderbook-4-building-the-order-book-cc3abb17c1ee)
- Building a Matching Engine : [Chapter 5](https://andriantriputra.medium.com)
- Market Orders vs Limit Orders : [Chapter 6](https://andriantriputra.medium.com)
- Order Cancellation : [Chapter 7](https://andriantriputra.medium.com)

## Phase 3 : Reliability

- Event-Sourced Order Book : [Chapter 8](https://andriantriputra.medium.com)
- Snapshot and Recovery : [Chapter 9](https://andriantriputra.medium.com)

## Phase 4 : Distribution

- Distributed Matching Engine : [Chapter 10](https://andriantriputra.medium.com)
- Market Sharding : [Chapter 11](https://andriantriputra.medium.com)
- Shard Ownership : [Chapter 12](https://andriantriputra.medium.com)
- Replicated Order Book : [Chapter 13](https://andriantriputra.medium.com)
- Fault-Tolerant Exchange : [Chapter 14](https://andriantriputra.medium.com)

## Phase 5 : Read Models

- CQRS for Market Data : [Chapter 15](https://andriantriputra.medium.com)

---

# 🚀 Running an Example

Every chapter is an independent Go module.

```bash
cd chapter-15
go run .
```

or choose any chapter you'd like to explore.

---

# 🎯 Who Is This For?

This series is designed for developers who want to understand how modern exchanges work internally, including:

- Backend Engineers
- Go Developers
- Software Engineers
- Distributed Systems Engineers
- System Design learners

No prior knowledge of trading systems is required.

---

# 👤 Author

**Andrian Tri Putra**

- Medium: https://andriantriputra.medium.com
- GitHub: https://github.com/andriantp
- GitHub (alternative): https://github.com/AndrianTriPutra

---

## 📄 License

This project is licensed under the MIT License.

---

If you find this repository useful, consider giving it a ⭐ on GitHub.