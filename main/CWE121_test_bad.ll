; ModuleID = 'CWE121_test_bad.c'
source_filename = "CWE121_test_bad.c"
target datalayout = "e-m:o-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-apple-macosx10.15.0"

@__const.CWE121_Stack_Based_Buffer_Overflow__CWE193_char_alloca_loop_01_bad.source = private unnamed_addr constant [11 x i8] c"AAAAAAAAAA\00", align 1

; Function Attrs: noinline nounwind optnone ssp uwtable
define void @CWE121_Stack_Based_Buffer_Overflow__CWE193_char_alloca_loop_01_bad() #0 {
  %1 = alloca i8*, align 8
  %2 = alloca i8*, align 8
  %3 = alloca [11 x i8], align 1
  %4 = alloca i64, align 8
  %5 = alloca i64, align 8
  %6 = alloca i8, i64 10, align 16
  store i8* %6, i8** %2, align 8
  %7 = load i8*, i8** %2, align 8
  store i8* %7, i8** %1, align 8
  %8 = load i8*, i8** %1, align 8
  %9 = getelementptr inbounds i8, i8* %8, i64 0
  store i8 0, i8* %9, align 1
  %10 = bitcast [11 x i8]* %3 to i8*
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 1 %10, i8* align 1 getelementptr inbounds ([11 x i8], [11 x i8]* @__const.CWE121_Stack_Based_Buffer_Overflow__CWE193_char_alloca_loop_01_bad.source, i32 0, i32 0), i64 11, i1 false)
  %11 = getelementptr inbounds [11 x i8], [11 x i8]* %3, i64 0, i64 0
  %12 = call i64 @strlen(i8* %11)
  store i64 %12, i64* %5, align 8
  store i64 0, i64* %4, align 8
  br label %13

13:                                               ; preds = %25, %0
  %14 = load i64, i64* %4, align 8
  %15 = load i64, i64* %5, align 8
  %16 = add i64 %15, 1
  %17 = icmp ult i64 %14, %16
  br i1 %17, label %18, label %28

18:                                               ; preds = %13
  %19 = load i64, i64* %4, align 8
  %20 = getelementptr inbounds [11 x i8], [11 x i8]* %3, i64 0, i64 %19
  %21 = load i8, i8* %20, align 1
  %22 = load i8*, i8** %1, align 8
  %23 = load i64, i64* %4, align 8
  %24 = getelementptr inbounds i8, i8* %22, i64 %23
  store i8 %21, i8* %24, align 1
  br label %25

25:                                               ; preds = %18
  %26 = load i64, i64* %4, align 8
  %27 = add i64 %26, 1
  store i64 %27, i64* %4, align 8
  br label %13

28:                                               ; preds = %13
  %29 = load i8*, i8** %1, align 8
  call void @printLine(i8* %29)
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
