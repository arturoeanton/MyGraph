# MyGraph (Academic/Experimental)

This project is an **academic implementation** of a Memory Graph Database using **Go** and **Antlr4**. It features a custom set implementation inspired by the principles of **CONCISE** (Compressed 'n' Composable Integer Set), but diverges in its architectural approach to explore specific trade-offs in sparse data management.

> **Note:** This project is intended for educational and research purposes. For production systems requiring high-performance bitmap compression, it is recommended to use **Roaring Bitmaps** (see *Industry Standard* section below).

## Theoretical Architecture: Hash-based Sparse BitSet

Unlike the standard CONCISE algorithm, which uses **Run-Length Encoding (RLE)** over sequential arrays to compress data, this project implements a **Hash-based Sparse BitSet** using Go's `map[uint]uint64`.

### Comparative Analysis: Hash Map vs. RLE (CONCISE)

The core difference lies in how the algorithms handle the **Universe of Discourse** (the range of possible integer values):

1.  **Space Complexity & Sparsity:**
    * **CONCISE (RLE):** Excellent for *dense* data or data with long consecutive runs of 0s or 1s. It compresses these runs into single words. However, for highly random/dispersed data (high entropy), the compression gain is negligible or negative due to overhead.
    * **MyGraph (Hash-based):** Optimized for **extreme sparsity**. It only allocates memory for 64-bit blocks that contain at least one set bit. Empty space costs **0 bytes**.
        * *Theoretical Implication:* If the density $d \to 0$, the space complexity approximates $O(k)$ where $k$ is the number of set blocks, whereas RLE approaches might incur $O(N)$ overhead to encode empty gaps if not optimized efficiently.

2.  **Time Complexity (Random Access):**
    * **CONCISE:** To check if a bit exists (`Contains`), RLE algorithms typically require **$O(N)$** or **$O(\log N)$** time (depending on auxiliary indexing), as they must decode or traverse the compressed sequence to find the target position.
    * **MyGraph:** By leveraging the hash map, this implementation achieves amortized **$O(1)$** random access. This is theoretically superior for workloads dominated by random lookups rather than sequential scans.

## Industry Standard: Why Roaring Bitmaps?

For most modern production scenarios, **[Roaring Bitmaps](https://roaringbitmap.org/)** are considered the state-of-the-art solution, surpassing both straightforward RLE (like CONCISE/WAH) and pure Hash-based implementations.

**Theoretical Justification for Roaring:**
Roaring employs a **hybrid dynamic container** strategy. It partitions the data space into chunks (e.g., $2^{16}$ integers) and dynamically decides the storage strategy for *each chunk* based on its cardinality:
* **Sparse Chunks:** Stored as simple Arrays of integers (low overhead, fast iteration).
* **Dense Chunks:** Stored as uncompressed Bitmaps (fast CPU bitwise operations).
* **Run Chunks:** Stored using RLE (compression for consecutive data).

**Conclusion:**
While **MyGraph** offers a specialized optimization for *random access in sparse datasets* via hashing, **Roaring Bitmaps** provides the best "worst-case" guarantee and overall performance stability by adapting its internal structure to the data distribution, avoiding the memory overhead of Go maps in dense scenarios and the scan latency of RLE in sparse ones.
