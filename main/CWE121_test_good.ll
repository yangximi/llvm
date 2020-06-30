; ModuleID = 'CWE121_test_good.c'
source_filename = "CWE121_test_good.c"
target datalayout = "e-m:o-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-apple-macosx10.15.0"

@__const.CWE121_Stack_Based_Buffer_Overflow__CWE193_char_alloca_loop_01_good.source = private unnamed_addr constant [11 x i8] c"AAAAAAAAAA\00", align 1

; Function Attrs: noinline nounwind optnone ssp uwtable
define void @CWE121_Stack_Based_Buffer_Overflow__CWE193_char_alloca_loop_01_good() #0 {
  %1 = alloca i8*, align 8
  %2 = alloca i8*, align 8
  %3 = alloca i8*, align 8
  %4 = alloca [11 x i8], align 1
  %5 = alloca i64, align 8
  %6 = alloca i64, align 8
  %7 = alloca i8, i64 11, align 16
  store i8* %7, i8** %2, align 8
  %8 = alloca i8, i64 11, align 16
  store i8* %8, i8** %3, align 8
  %9 = load i8*, i8** %2, align 8
  store i8* %9, i8** %1, align 8
  %10 = load i8*, i8** %1, align 8
  %11 = getelementptr inbounds i8, i8* %10, i64 0
  store i8 0, i8* %11, align 1
  %12 = bitcast [11 x i8]* %4 to i8*
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 1 %12, i8* align 1 getelementptr inbounds ([11 x i8], [11 x i8]* @__const.CWE121_Stack_Based_Buffer_Overflow__CWE193_char_alloca_loop_01_good.source, i32 0, i32 0), i64 11, i1 false)
  %13 = getelementptr inbounds [11 x i8], [11 x i8]* %4, i64 0, i64 0
  %14 = call i64 @strlen(i8* %13)
  store i64 %14, i64* %6, align 8
  store i64 0, i64* %5, align 8
  br label %15

15:                                               ; preds = %27, %0
  %16 = load i64, i64* %5, align 8
  %17 = load i64, i64* %6, align 8
  %18 = add i64 %17, 1
  %19 = icmp ult i64 %16, %18
  br i1 %19, label %20, label %30

20:                                               ; preds = %15
  %21 = load i64, i64* %5, align 8
  %22 = getelementptr inbounds [11 x i8], [11 x i8]* %4, i64 0, i64 %21
  %23 = load i8, i8* %22, align 1
  %24 = load i8*, i8** %1, align 8
  %25 = load i64, i64* %5, align 8
  %26 = getelementptr inbounds i8, i8* %24, i64 %25
  store i8 %23, i8* %26, align 1
  br label %27

27:                                               ; preds = %20
  %28 = load i64, i64* %5, align 8
  %29 = add i64 %28, 1
  store i64 %29, i64* %5, align 8
  br label %15

30:                                               ; preds = %15
  %31 = load i8*, i8** %1, align 8
  call void @printLine(i8* %31)
  ret void
}

; Function Attrs: argmemonly nounwind
declare void @llvm.memcpy.p0i8.p0i8.i64(i8* nocapture writeonly, i8* nocapture readonly, i64, i1 immarg) #1

declare i64 @strlen(i8*) #2

declare void @printLine(i8*) #2

attributes #0 = { noinline nounwind optnone ssp uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "darwin-stkchk-strong-link" "disable-tail-calls"="false" "frame-pointer"="all" "less-precise-fpmad"="false" "min-legal-vector-width"="0" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "probe-stack"="___chkstk_darwin" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+cx8,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { argmemonly nounwind }
attributes #2 = { "correctly-rounded-divide-sqrt-fp-math"="false" "darwin-stkchk-strong-link" "disable-tail-calls"="false" "frame-pointer"="all" "less-precise-fpmad"="false" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "probe-stack"="___chkstk_darwin" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+cx8,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.module.flags = !{!0, !1, !2}
!llvm.ident = !{!3}

!0 = !{i32 2, !"SDK Version", [3 x i32] [i32 10, i32 15, i32 4]}
!1 = !{i32 1, !"wchar_size", i32 4}
!2 = !{i32 7, !"PIC Level", i32 2}
!3 = !{!"Apple clang version 11.0.3 (clang-1103.0.32.29)"}
