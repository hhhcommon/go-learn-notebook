/**
 * 算法 2.1
 * 
 * 将所有在Lb 不在La 的元素合并到La中
 * 
 */
void union(List &La, List Lb)
{
    La_len = ListLength(La); // 求长度
    Lb_len = ListLength(Lb);

    for (i = 0; i <= Lb_len; i++) // 遍历Lb元素
    {
        GetElem(Lb, i, e);             // 取出Lb的索引为i的元素赋值给e
        if (!LocateElem(La, e, equal)) // 判断La是否存在相同的元素
        {
            ListInsert(La, ++La_len);
        }
    }
}

/*
算法 2.2

合并两个线性表。

奇葩要求是，非递减排列。。。。。

处理这个问题的方式竟然是：
1. 初始化新的表 C
2. 同时遍历AB两个表，先把比较小的元素放在C，这是肯定会拿空一个表
3. 然后针对没有拿空的那个表继续挨个拿出来放到C。

这么做就能达到非递减（可能递增 有可能无序）

已知线性表La和Lb中的数据元素按值非递减排列。
归并La和Lb得到新的先行表Lc，
Lc的数据元素也按值非递减排列。
*/
void MergeList(List La, List Lb, List &Lc)
{
    InitList(Lc);
    i = j = 1;
    k = 0;
    La_len = ListLength(La);
    Lb_len = ListLength(Lb);

    // La 和 Lb 均非空 ，
    // 这时比较拿出来的元素，把小的先放到Lc
    // 最终Lb和La会先拿空一个
    while ((i <= La_len) && (j <= Lb_len))
    {
        GetElem(La, i, ai);
        GetElem(Lb, j, bj);
        if (ai < bj)
        {
            ListInsert(Lc, ++k, ai);
            ++i;
        }
        else
        {
            ListInsert(Lc, ++k, bj);
            ++j;
        }
    }

    // 针对没有拿干净A的情况
    while (i <= La_len)
    {
        GetElem(La, i++, ai);
        ListInsert(Lc, ++k, ai);
    }

    // 针对没有拿干净b的情况
    while (j <= Lb_len)
    {
        GetElem(La, j++, bj);
        ListInsert(Lc, ++k, bj);
    }
}

#include <stdlib.h>
#include <stdio.h>
// 线性表动态分配顺序存储结构
#define LIST_INT_SIZE 100 // 线性表存储空间初始化分配量
#define LISTINCREMENT 10  // 线性表存储空间的分配增量
typedef struct
{
    Elemtype *elem; // 存储空间基址
    int length;     // 当前长度
    int listsize;   // 当前分配的存储容量（以 sizeof(Elemtype)为单位）
} Sqlist;

/**
 * 算法 2.3
 * 
 * 构造一个空的线性表L
 * 
 * 方法：
 * 直接从内存分配一段长度。
 */
Status InitList_Sq(Sqlist &L){
    L.elem = (ElemType *)malloc(LIST_INT_SIZE * sizeof(ElemType))

}

/**
 * 算法 2.4 
 * 
 * 在顺序线性表L中第i个位置之前插入新的元素e
 * 
 * 解释：
 * 一般情况下，在第 i(1 <= i <= n)个元素之前插入一个元素时，
 * 需要将第n至第i个元素（共 n - i + 1 个元素）向后移动一个位置。
 * 
 */
Status ListInsert_Sq(Sqlist &L, int i, ElemType e)
{
    // 判断i的合法值
    // i的合法值是 1 <= i <= List_Length(L) + 1
    if (i < 1 || i > L.length + 1)
    {
        return ERROR; // i值不合法
    }

    if (L.length >= L.listsize)
    {
        // 当分配空间已满，重新按增量规则增加空间
        newbase = (ElemType *)realloc(L.elem,
                                      (L.listsize + LISTINCREMENT) * sizeof(ElemType));

        // 空间增加失败
        if (!newbase)
        {
            exist(OVERFLOW);
        }

        L.elem = newbase;            // 重新定义基址
        L.listsize += LISTINCREMENT; // 重新定义大小
    }

    q = &(L.elem[i - 1]); // 这是要插入的位置i-1 的 指针？
    for (p = &(L.elem(L.length - 1)); p >= q; --p)
    {
        // 遍历 i -》 n 的元素 更改他们的指针地址。
        *(p + 1) = *p
    }
    *q = e;     // 插入e
    ++L.length; //表长增长1
    return OK;
}

/**
 * 算法 2.5
 * 
 * 在顺序线性表L中删除第i个元素，并用e返回这个元素的值。
 * 
 */

Status ListDelet_Sq(SqList &L, int i, ElemType &e)
{
    if ((i < 1) || (i > L.length))
    {
        return ERROR;
    }

    p = &(L.elem[i - 1]);      // 头
    e = *p;                    // 被删除的元素
    q = L.elem + L.length - 1; // 尾
    for (++p, p <= q; ++p)
    {
        *(p - 1) = *p; // 被删元素之后的元素左移
    }
    --L.length; // 表长度减1
    return OK;
}

/**
 * 算法 2.6 
 * 
 * 在顺序线性表L中查找第一个与e满足Compare()的元素的位序。
 * 若找到，则返回其在L中的位序，否则返回0.
 * 
 */
int LocateElem_Sq(SqList L, ElemType e, Status (*compare)(ElemType, ElemType))
{
    i = 1;
    p = L.elem;
    while (i <= L.length && !(*compare)(*p++, e))
    {
        ++i;
    }
    if (i <= L.length)
    {
        return i;
    }
    else
    {
        return 0;
    }
}

/**
 * 算法 2.7  
 * 
 * 已知顺序线性表La 和 Lb 的元素按值非递减排列。
 * 归并La和Lb的到新的顺序线性表Lc，Lc也按值非递减排列。
 * 
 */

void MergeList_Sq(SqList La, SqList Lb, SqList &Lc)
{
    pa = La.elem;
    pb = Lb.elem;

    Lc.listsize = Lc.length = La.length + Lb.length;

    pc = Lc.elem = (ElemType *)malloc(Lc.listsize * sizeof(ElemType));

    if (!Lc.elem)
    {
        exist(OVVERFLOW);
    }

    pa_last = La.elem + La.length - 1;
    pb_last = Lb.elem + Lb.length - 1;

    while (pa <= pa_last && pb <= pb_last)
    {
        if (*pa <= *pb)
        {
            *pc++ = *pa++;
        }
        else
        {
            *pc++ = *pb++
        }
    }

    while (pa <= pa_last)
    {
        *pc++ = *pa++;
    }
    while (pb <= pb_last)
    {
        *pc++ = *pb++;
    }
}