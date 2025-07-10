### LFU (Least Frequently Used) Cache

**LFU** is a cache eviction strategy that removes the least frequently accessed items, rather than just the oldest unused ones.

LFU is effective in systems where long-term data popularity matters. For example, in an e-commerce app, frequently viewed or purchased products stay in cache, while rarely accessed items get removed.  

Although LFU is more complex because it requires tracking access counts, it typically results in a lower cache miss rate over time.

---

### LFU (Least Frequently Used) Cache

**LFU** — это стратегия кэширования, при которой удаляются элементы, к которым обращались реже всего, а не просто те, что давно не использовались.

LFU хорошо работает в системах, где важна долгосрочная востребованность данных — например, при кэшировании товаров в интернет-магазине. Часто просматриваемые товары остаются в памяти, а те, что покупают раз в год – удаляются.

Алгоритм сложнее, так как требует учета количества обращений, но в долгосрочной перспективе он снижает количество промахов кэша.
