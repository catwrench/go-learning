内容说明：

- 《算法4》的 golang 版本实现 （部分使用了泛型 ~~自作自受orz~~）
    - java参考：https://algs4.cs.princeton.edu/
    - golang参考：https://github.com/hashyong/algorithm4
- 常用数据结构
- 日常随笔

施工进度......
***

- datastructures 数据结构
    - bag 背包
        1.[x] Bag : 背包
    - channel 管道
        1.[x] PushSpeedControl : 动态控制发送速率
    - queue 队列
        1. [x] Queue : 队列（基于链表）
        2. [x] MaxPQ : 最大优先队列(基于堆)
        3. [x] MinPQ : 最小优先队列(基于堆)
        4. [x] IndexMinPQ : 索引最小优先队列
    - stack 栈
        1. [x] LIFOStack : 后进先出栈
    - st 符号表
        1. [x] SequentialSearchST : 顺序查找，基于链表
        2. [x] BinarySearchST : 二分查找，基于数组
        3. [x] BST : 二叉查找树
        4. [x] RedBlackBST : 红黑树
        5. [x] SeparateChainingHashST : 散列表，拉链法
        6. [x] LinearProbingHashST : 散列表，线性探测
    - graph 图
        1. [x] graph : 无向图，邻接表实现
        2. [x] SymbolGraph : 符号图
        3. [x] DiGraph : 有向图
        4. [x] Edge : 带权重的边
        5. [x] EdgeWeightedGraph : 加权无向图
        6. [x] DirectedEdge 带权重的有向边
        7. [x] EdgeWeightedDiGraph : 加权有向图
    - others 其他
        1. [x] LookUpIndex : 索引（及反向索引）查找, 基于 SeparateChainingHashST
        2. [x] SparseVector : 稀疏向量, 基于 SeparateChainingHashST
        3. [x] SymbolIndexMinPQ : 符号索引最小优先队列

***

- algorithm 算法
    - search 搜索
        - simpleSearch : 简单查找
            1. [x] BinarySearch : 二分查找
    - sort 排序
        1. [x] SortCompare : 排序算法性能对比
        2. [x] SelectionSort : 选择排序
        3. [x] InsertionSort : 插入排序

        - quickSort 快速排序
            1. [x] QuickSort : 快速排序
            2. [x] QuickSort3Way : 三向切分的快速排序

        4.[x] ShellSort : 希尔排序

        - MergeSort 归并排序
            1. [x] MergeUBSort : 自顶向下的归并
            2. [x] MergeBUSort : 自底向上的归并

        5.[x] HeapSort : 堆排序
    - string 字符串查找/排序
        1. [x] CountIndex : 键索引记数法排序
        2. [x] LSD : 低位优先排序
        3. [x] MSD : 高位优先排序
        4. [x] Quick3String : 三向切分的字符串排序
        5. [x] TrieST : 基于单词查找树的符号表
        6. [x] TST : 三向单词查找树
        7. [x] KMP : 字符串的搜索
        8. [x] BoyerMoore : 博伊尔-摩尔 字符串搜索(启发式的处理不匹配的字符)
        9. [x] RabinKarp : 指纹(散列)字符串查找算法
        10. [ ] Huffman : 霍夫曼压缩

    - tree 树
        1. [x] tree : 遍历
    - graph 图
        - unionFind 动态连通性算法
            1. [x] QuickFind : 连通分量查找，QuickFind算法
            2. [x] QuickUnion : 连通分量查找，QuickUnion加权算法
        - dfs 深度搜索优先（无向图）
            1. [x] dfs : 单点连通性，图的深度搜索优先搜索
            2. [x] DepthFirstPaths : 单点路径
            3. [x] CC : 连通分量查找
            4. [x] Cycle : 检测是否存在环
            5. [x] TwoColor : 双色问题（图的二分性）
        - ddfs 深度优先搜索（有向图）
            1. [x] DirectedDFS : 单点多点可达性
            2. [x] DirectedCycle : 有向环检测
            3. [x] DepthFirstOrder : 基于深度优先搜索的顶点排序
            4. [x] Topological : 拓扑排序，优先级限制下的调度问题
            5. [x] KosarajuSCC : 强连通性
            6. [x] TransitiveClosure : 传递闭包（可到达性，类比无向图的单点路径）
        - bfs 广度优先搜索
            1. [x] BreadthFirstPaths : 广度搜索优先，路径查找
            2. [x] DegreesOfSeparation : 分隔的度
        - mst 最小生成树
            1. [x] LazyPrimMST : 最小生成树-prim算法的延时实现
            2. [x] PrimMST : 最小生成树-即时普利姆算法
            3. [x] KruskalMST : 最小生成树-克鲁斯卡尔算法
        - sp 单点路径
            1. [x] DijkstraSP : 单点最短路径-迪杰斯特拉算法
            2. [x] AcyclicSP : 无环加权有向图的最短路径算法
            3. [x] AcyclicLP : 无环加权有向图的最长路径算法
            4. [x] CMP : 关键路径 优先级限制下的并行任务调度问题
            5. [x] BellManFordSP : 基于队列的BellMan-Ford算法
            6. [x] Arbitrage : 套汇获利问题

***

- concurrency 并发
    - concurrency : SyncWait 并发
    - goroutine : WaitGroup

