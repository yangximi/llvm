; ModuleID = 'CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad.cpp'
source_filename = "CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad.cpp"
target datalayout = "e-m:o-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-apple-macosx10.15.0"

%"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad" = type { i8* }

@.str = private unnamed_addr constant [43 x i32] [i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 65, i32 0], align 4

; Function Attrs: noinline nounwind optnone ssp uwtable
define void @_ZN45CWE121_Stack_Based_Buffer_Overflow__CWE135_8449CWE121_Stack_Based_Buffer_Overflow__CWE135_84_badC2EPv(%"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, i8*) unnamed_addr #0 align 2 {
  %3 = alloca %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, align 8
  %4 = alloca i8*, align 8
  store %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %0, %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"** %3, align 8
  store i8* %1, i8** %4, align 8
  %5 = load %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"** %3, align 8
  %6 = load i8*, i8** %4, align 8
  %7 = getelementptr inbounds %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad", %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %5, i32 0, i32 0
  store i8* %6, i8** %7, align 8
  %8 = getelementptr inbounds %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad", %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %5, i32 0, i32 0
  store i8* bitcast ([43 x i32]* @.str to i8*), i8** %8, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone ssp uwtable
define void @_ZN45CWE121_Stack_Based_Buffer_Overflow__CWE135_8449CWE121_Stack_Based_Buffer_Overflow__CWE135_84_badC1EPv(%"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, i8*) unnamed_addr #0 align 2 {
  %3 = alloca %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, align 8
  %4 = alloca i8*, align 8
  store %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %0, %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"** %3, align 8
  store i8* %1, i8** %4, align 8
  %5 = load %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"** %3, align 8
  %6 = load i8*, i8** %4, align 8
  call void @_ZN45CWE121_Stack_Based_Buffer_Overflow__CWE135_8449CWE121_Stack_Based_Buffer_Overflow__CWE135_84_badC2EPv(%"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %5, i8* %6)
  ret void
}

; Function Attrs: noinline optnone ssp uwtable
define void @_ZN45CWE121_Stack_Based_Buffer_Overflow__CWE135_8449CWE121_Stack_Based_Buffer_Overflow__CWE135_84_badD2Ev(%"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*) unnamed_addr #1 align 2 {
  %2 = alloca %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, align 8
  %3 = alloca i64, align 8
  %4 = alloca i8*, align 8
  store %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %0, %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"** %2, align 8
  %5 = load %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"** %2, align 8
  %6 = getelementptr inbounds %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad", %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %5, i32 0, i32 0
  %7 = load i8*, i8** %6, align 8
  %8 = call i64 @strlen(i8* %7)
  store i64 %8, i64* %3, align 8
  %9 = load i64, i64* %3, align 8
  %10 = add i64 %9, 1
  %11 = mul i64 %10, 4
  %12 = alloca i8, i64 %11, align 16
  store i8* %12, i8** %4, align 8
  %13 = load i8*, i8** %4, align 8
  %14 = bitcast i8* %13 to i32*
  %15 = getelementptr inbounds %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad", %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %5, i32 0, i32 0
  %16 = load i8*, i8** %15, align 8
  %17 = bitcast i8* %16 to i32*
  %18 = call i32* @wcscpy(i32* %14, i32* %17)
  %19 = load i8*, i8** %4, align 8
  call void @printLine(i8* %19)
  ret void
}

declare i64 @strlen(i8*) #2

declare i32* @wcscpy(i32*, i32*) #2

declare void @printLine(i8*) #2

; Function Attrs: noinline optnone ssp uwtable
define void @_ZN45CWE121_Stack_Based_Buffer_Overflow__CWE135_8449CWE121_Stack_Based_Buffer_Overflow__CWE135_84_badD1Ev(%"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*) unnamed_addr #1 align 2 {
  %2 = alloca %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, align 8
  store %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %0, %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"** %2, align 8
  %3 = load %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"*, %"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"** %2, align 8
  call void @_ZN45CWE121_Stack_Based_Buffer_Overflow__CWE135_8449CWE121_Stack_Based_Buffer_Overflow__CWE135_84_badD2Ev(%"class.CWE121_Stack_Based_Buffer_Overflow__CWE135_84::CWE121_Stack_Based_Buffer_Overflow__CWE135_84_bad"* %3)
  ret void
}

attributes #0 = { noinline nounwind optnone ssp uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "darwin-stkchk-strong-link" "disable-tail-calls"="false" "frame-pointer"="all" "less-precise-fpmad"="false" "min-legal-vector-width"="0" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "probe-stack"="___chkstk_darwin" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+cx8,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { noinline optnone ssp uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "darwin-stkchk-strong-link" "disable-tail-calls"="false" "frame-pointer"="all" "less-precise-fpmad"="false" "min-legal-vector-width"="0" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "probe-stack"="___chkstk_darwin" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+cx8,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #2 = { "correctly-rounded-divide-sqrt-fp-math"="false" "darwin-stkchk-strong-link" "disable-tail-calls"="false" "frame-pointer"="all" "less-precise-fpmad"="false" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "probe-stack"="___chkstk_darwin" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+cx8,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.module.flags = !{!0, !1, !2}
!llvm.ident = !{!3}

!0 = !{i32 2, !"SDK Version", [3 x i32] [i32 10, i32 15, i32 4]}
!1 = !{i32 1, !"wchar_size", i32 4}
!2 = !{i32 7, !"PIC Level", i32 2}
!3 = !{!"Apple clang version 11.0.3 (clang-1103.0.32.29)"}
