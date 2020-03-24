// page.10

// 函数结果状态码
#define  TRUE 1
#define  FALSE 0
#define  OK 1
#define  ERROR 0
#define  INFEASIBLE -1
#define  OVERFLOW -2

typedef int ElemType;

// Status 是函数类型，其值是函数结果状态码
typedef int Status;

typedef ElemType *Triplet;

// 函数声明部分

Status InitTriplet(Triplet &T, ElemType v1, ElemType v2, ElemType v3);

Status DestroyTriplet(Triplet &T);

Status Get(Triplet T, int i, ElemType &e);

Status Put(Triplet &T, int i, ElemType e);

Status IsAcending(Triplet T);

Status IsDescending(Triplet T);

Status Max(Triplet T, ElemType &e);

Status Min(Triplet T, ElemType &e);
