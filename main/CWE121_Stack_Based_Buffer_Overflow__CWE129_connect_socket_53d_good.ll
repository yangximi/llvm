source_filename = "CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_53d.c"
target datalayout = "e-m:o-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-apple-macosx10.15.0"

@.str = private unnamed_addr constant [32 x i8] c"ERROR: Array index is negative.\00", align 1
@.str.1 = private unnamed_addr constant [36 x i8] c"ERROR: Array index is out-of-bounds\00", align 1

define void @CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_53d_goodB2GSink(i32 %0) #0 {
1:
	%2 = alloca i32, align 4
	%3 = alloca i32, align 4
	%4 = alloca [10 x i32], align 16
	store i32 %0, i32* %2, align 4
	%5 = bitcast [10 x i32]* %4 to i8*
	call void @llvm.memset.p0i8.i64(i8* align 16 %5, i8 0, i64 40, i1 false)
	%6 = load i32, i32* %2, align 4
	%7 = icmp sge i32 %6, 0
	br i1 %7, label %8, label %27

8:
	%9 = load i32, i32* %2, align 4
	%10 = icmp slt i32 %9, 10
	br i1 %10, label %11, label %27

11:
	%12 = load i32, i32* %2, align 4
	%13 = sext i32 %12 to i64
	%14 = getelementptr inbounds [10 x i32], [10 x i32]* %4, i64 0, i64 %13
	store i32 1, i32* %14, align 4
	store i32 0, i32* %3, align 4
	br label %15

15:
	%16 = load i32, i32* %3, align 4
	%17 = icmp slt i32 %16, 10
	br i1 %17, label %18, label %26

18:
	%19 = load i32, i32* %3, align 4
	%20 = sext i32 %19 to i64
	%21 = getelementptr inbounds [10 x i32], [10 x i32]* %4, i64 0, i64 %20
	%22 = load i32, i32* %21, align 4
	call void @printIntLine(i32 %22)
	br label %23

23:
	%24 = load i32, i32* %3, align 4
	%25 = add nsw i32 %24, 1
	store i32 %25, i32* %3, align 4
	br label %15

26:
	br label %28

27:
	call void @printLine(i8* getelementptr inbounds ([36 x i8], [36 x i8]* @.str.1, i64 0, i64 0))
	br label %28

28:
	ret void
}

declare void @llvm.memset.p0i8.i64(i8* nocapture writeonly %0, i8 %1, i64 %2, i1 immarg %3) #1

declare void @printIntLine(i32 %0) #2

declare void @printLine(i8* %0) #2

attributes #0 = { noinline nounwind optnone ssp uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "darwin-stkchk-strong-link" "disable-tail-calls"="false" "frame-pointer"="all" "less-precise-fpmad"="false" "min-legal-vector-width"="0" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "probe-stack"="___chkstk_darwin" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+cx8,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { argmemonly nounwind }
attributes #2 = { "correctly-rounded-divide-sqrt-fp-math"="false" "darwin-stkchk-strong-link" "disable-tail-calls"="false" "frame-pointer"="all" "less-precise-fpmad"="false" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "probe-stack"="___chkstk_darwin" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+cx8,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!3}
!llvm.module.flags = !{!0, !1, !2}

!0 = !{i32 2, !"SDK Version", [3 x i32] [i32 10, i32 15, i32 4]}
!1 = !{i32 1, !"wchar_size", i32 4}
!2 = !{i32 7, !"PIC Level", i32 2}
!3 = !{!"Apple clang version 11.0.3 (clang-1103.0.32.29)"}

