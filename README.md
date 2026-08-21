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
- Building a Matching Engine : [Chapter 5](https://medium.com/@andriantriputra/orderbook-5-building-a-matching-engine-aca5fbc64742)
- Market Orders vs Limit Orders : [Chapter 6](https://medium.com/@andriantriputra/orderbook-6-market-orders-vs-limit-orders-c5e09185c2f8)
- Order Cancellation : [Chapter 7](https://medium.com/@andriantriputra/orderbook-7-order-cancellation-b7d4205e3eac)

## Phase 3 : Reliability

- Event-Sourced Order Book : [Chapter 8](https://medium.com/@andriantriputra/orderbook-8-event-sourced-order-book-778dc8b14483)
- Snapshot and Recovery : [Chapter 9](https://medium.com/@andriantriputra/orderbook-9-snapshot-and-recovery-155b0ff133a3)

## Phase 4 : Distribution

- Distributed Matching Engine : [Chapter 10](https://medium.com/@andriantriputra/orderbook-10-distributed-matching-engine-1da770164b1e)
- Market Sharding : [Chapter 11](https://medium.com/@andriantriputra/orderbook-11-market-sharding-69bbebe46651)
- Shard Ownership : [Chapter 12](https://medium.com/@andriantriputra/orderbook-12-shard-ownership-57fa9f98f751)
- Replicated Order Book : [Chapter 13](https://medium.com/@andriantriputra/orderbook-13-replicated-order-book-c53a1728cf82)
- Fault-Tolerant Exchange : [Chapter 14](https://medium.com/@andriantriputra/orderbook-14-fault-tolerant-exchange-7af757c9056d)

## Phase 5 : Read Models

- CQRS for Market Data : [Chapter 15](https://medium.com/@andriantriputra/orderbook-15-cqrs-for-market-data-ca664d09e969)

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