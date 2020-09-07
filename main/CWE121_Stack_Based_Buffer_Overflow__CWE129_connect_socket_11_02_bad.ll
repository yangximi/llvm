source_filename = "CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_11.c"
target datalayout = "e-m:o-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-apple-macosx10.15.0"

%struct.in_addr = type { i32 }
%struct.sockaddr = type { i8, i8, [14 x i8] }
%struct.sockaddr_in = type { i8, i8, i16, %struct.in_addr, [8 x i8] }

@.str = private unnamed_addr constant [10 x i8] c"127.0.0.1\00", align 1
@.str.1 = private unnamed_addr constant [32 x i8] c"ERROR: Array index is negative.\00", align 1
@.str.2 = private unnamed_addr constant [18 x i8] c"Calling good()...\00", align 1
@.str.3 = private unnamed_addr constant [16 x i8] c"Finished good()\00", align 1
@.str.4 = private unnamed_addr constant [17 x i8] c"Calling bad()...\00", align 1
@.str.5 = private unnamed_addr constant [15 x i8] c"Finished bad()\00", align 1
@.str.6 = private unnamed_addr constant [21 x i8] c"Benign, fixed string\00", align 1
@.str.7 = private unnamed_addr constant [36 x i8] c"ERROR: Array index is out-of-bounds\00", align 1

define void @CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_11_bad() #0 !dbg !28 {
0:
	%1 = alloca i32, align 4
	%2 = alloca i32, align 4
	%3 = alloca %struct.sockaddr_in, align 4
	%4 = alloca i32, align 4
	%5 = alloca [14 x i8], align 1
	%6 = alloca i32, align 4
	%7 = alloca [10 x i32], align 16
	call void @llvm.dbg.declare(metadata i32* %1, metadata !31, metadata !DIExpression()), !dbg !33
	store i32 -1, i32* %1, align 4, !dbg !34
	%8 = call i32 (...) @globalReturnsTrue(), !dbg !35
	%9 = icmp ne i32 %8, 0, !dbg !35
	br i1 %9, label %10, label %52, !dbg !37

10:
	call void @llvm.dbg.declare(metadata i32* %2, metadata !38, metadata !DIExpression()), !dbg !41
	call void @llvm.dbg.declare(metadata %struct.sockaddr_in* %3, metadata !42, metadata !DIExpression()), !dbg !62
	call void @llvm.dbg.declare(metadata i32* %4, metadata !63, metadata !DIExpression()), !dbg !64
	store i32 -1, i32* %4, align 4, !dbg !64
	call void @llvm.dbg.declare(metadata [14 x i8]* %5, metadata !65, metadata !DIExpression()), !dbg !66
	br label %11, !dbg !67

11:
	%12 = call i32 @socket(i32 2, i32 1, i32 6), !dbg !68
	store i32 %12, i32* %4, align 4, !dbg !70
	%13 = load i32, i32* %4, align 4, !dbg !71
	%14 = icmp eq i32 %13, -1, !dbg !73
	br i1 %14, label %15, label %16, !dbg !74

15:
	br label %45, !dbg !75

16:
	%17 = bitcast %struct.sockaddr_in* %3 to i8*, !dbg !77
	call void @llvm.memset.p0i8.i64(i8* align 4 %17, i8 0, i64 16, i1 false), !dbg !77
	%18 = getelementptr inbounds %struct.sockaddr_in, %struct.sockaddr_in* %3, i32 0, i32 1, !dbg !78
	store i8 2, i8* %18, align 1, !dbg !79
	%19 = call i32 @inet_addr(i8* getelementptr inbounds ([10 x i8], [10 x i8]* @.str, i64 0, i64 0)), !dbg !80
	%20 = getelementptr inbounds %struct.sockaddr_in, %struct.sockaddr_in* %3, i32 0, i32 3, !dbg !81
	%21 = getelementptr inbounds %struct.in_addr, %struct.in_addr* %20, i32 0, i32 0, !dbg !82
	store i32 %19, i32* %21, align 4, !dbg !83
	%22 = getelementptr inbounds %struct.sockaddr_in, %struct.sockaddr_in* %3, i32 0, i32 2, !dbg !84
	store i16 -30871, i16* %22, align 2, !dbg !85
	%23 = load i32, i32* %4, align 4, !dbg !86
	%24 = bitcast %struct.sockaddr_in* %3 to %struct.sockaddr*, !dbg !88
	%25 = call i32 @"\01_connect"(i32 %23, %struct.sockaddr* %24, i32 16), !dbg !89
	%26 = icmp eq i32 %25, -1, !dbg !90
	br i1 %26, label %27, label %28, !dbg !91

27:
	br label %45, !dbg !92

28:
	%29 = load i32, i32* %4, align 4, !dbg !94
	%30 = getelementptr inbounds [14 x i8], [14 x i8]* %5, i64 0, i64 0, !dbg !95
	%31 = call i64 @"\01_recv"(i32 %29, i8* %30, i64 13, i32 0), !dbg !96
	%32 = trunc i64 %31 to i32, !dbg !96
	store i32 %32, i32* %2, align 4, !dbg !97
	%33 = load i32, i32* %2, align 4, !dbg !98
	%34 = icmp eq i32 %33, -1, !dbg !100
	br i1 %34, label %38, label %35, !dbg !101

35:
	%36 = load i32, i32* %2, align 4, !dbg !102
	%37 = icmp eq i32 %36, 0, !dbg !103
	br i1 %37, label %38, label %39, !dbg !104

38:
	br label %45, !dbg !105

39:
	%40 = load i32, i32* %2, align 4, !dbg !107
	%41 = sext i32 %40 to i64, !dbg !108
	%42 = getelementptr inbounds [14 x i8], [14 x i8]* %5, i64 0, i64 %41, !dbg !108
	store i8 0, i8* %42, align 1, !dbg !109
	%43 = getelementptr inbounds [14 x i8], [14 x i8]* %5, i64 0, i64 0, !dbg !110
	%44 = call i32 @atoi(i8* %43), !dbg !111
	store i32 %44, i32* %1, align 4, !dbg !112
	br label %45, !dbg !113

45:
	%46 = load i32, i32* %4, align 4, !dbg !114
	%47 = icmp ne i32 %46, -1, !dbg !116
	br i1 %47, label %48, label %51, !dbg !117

48:
	%49 = load i32, i32* %4, align 4, !dbg !118
	%50 = call i32 @"\01_close"(i32 %49), !dbg !120
	br label %51, !dbg !121

51:
	br label %52, !dbg !122

52:
	%53 = call i32 (...) @globalReturnsTrue(), !dbg !123
	%54 = icmp ne i32 %53, 0, !dbg !123
	br i1 %54, label %55, label %77, !dbg !125

55:
	call void @llvm.dbg.declare(metadata i32* %6, metadata !126, metadata !DIExpression()), !dbg !129
	call void @llvm.dbg.declare(metadata [10 x i32]* %7, metadata !130, metadata !DIExpression()), !dbg !134
	%56 = bitcast [10 x i32]* %7 to i8*, !dbg !134
	call void @llvm.memset.p0i8.i64(i8* align 16 %56, i8 0, i64 40, i1 false), !dbg !134
	%57 = load i32, i32* %1, align 4, !dbg !135
	%58 = icmp sge i32 %57, 0, !dbg !137
	br i1 %58, label %59, label %75, !dbg !138

59:
	%60 = load i32, i32* %1, align 4, !dbg !139
	%61 = sext i32 %60 to i64, !dbg !141
	%62 = getelementptr inbounds [10 x i32], [10 x i32]* %7, i64 0, i64 %61, !dbg !141
	store i32 1, i32* %62, align 4, !dbg !142
	store i32 0, i32* %6, align 4, !dbg !143
	br label %63, !dbg !145

63:
	%64 = load i32, i32* %6, align 4, !dbg !146
	%65 = icmp slt i32 %64, 10, !dbg !148
	br i1 %65, label %66, label %74, !dbg !149

66:
	%67 = load i32, i32* %6, align 4, !dbg !150
	%68 = sext i32 %67 to i64, !dbg !152
	%69 = getelementptr inbounds [10 x i32], [10 x i32]* %7, i64 0, i64 %68, !dbg !152
	%70 = load i32, i32* %69, align 4, !dbg !152
	call void @printIntLine(i32 %70), !dbg !153
	br label %71, !dbg !154

71:
	%72 = load i32, i32* %6, align 4, !dbg !155
	%73 = add nsw i32 %72, 1, !dbg !155
	store i32 %73, i32* %6, align 4, !dbg !155
	br label %63, !dbg !156, !llvm.loop !157

74:
	br label %76, !dbg !159

75:
	call void @printLine(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.str.1, i64 0, i64 0)), !dbg !160
	br label %76

76:
	br label %77, !dbg !162

77:
	ret void, !dbg !163
}

declare void @llvm.dbg.declare(metadata %0, metadata %1, metadata %2) #1

declare i32 @globalReturnsTrue(...) #2

declare i32 @socket(i32 %0, i32 %1, i32 %2) #2

declare void @llvm.memset.p0i8.i64(i8* nocapture writeonly %0, i8 %1, i64 %2, i1 immarg %3) #3

declare i32 @inet_addr(i8* %0) #2

declare i32 @"\01_connect"(i32 %0, %struct.sockaddr* %1, i32 %2) #2

declare i64 @"\01_recv"(i32 %0, i8* %1, i64 %2, i32 %3) #2

declare i32 @atoi(i8* %0) #2

declare i32 @"\01_close"(i32 %0) #2

declare void @printIntLine(i32 %0) #2

declare void @printLine(i8* %0) #2

declare void @srand(i32 %0) #2

declare i64 @time(i64* %0) #2

declare i32 @globalReturnsFalse(...) #2

attributes #0 = { noinline nounwind optnone ssp uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "disable-tail-calls"="false" "frame-pointer"="all" "less-precise-fpmad"="false" "min-legal-vector-width"="0" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+cx8,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { nounwind readnone speculatable willreturn }
attributes #2 = { "correctly-rounded-divide-sqrt-fp-math"="false" "disable-tail-calls"="false" "frame-pointer"="all" "less-precise-fpmad"="false" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+cx8,+fxsr,+mmx,+sahf,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #3 = { argmemonly nounwind willreturn }

!llvm.dbg.cu = !{!0}
!llvm.ident = !{!27}
!llvm.module.flags = !{!23, !24, !25, !26}

!0 = distinct !DICompileUnit(language: DW_LANG_C99, file: !1, producer: "clang version 10.0.0 ", emissionKind: FullDebug, enums: !2, retainedTypes: !3, nameTableKind: None)
!1 = !DIFile(filename: "CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_11.c", directory: "/Users/ys/Project/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01")
!2 = !{}
!3 = !{!4, !7, !22}
!4 = !DIDerivedType(tag: DW_TAG_typedef, name: "__uint16_t", file: !5, line: 43, baseType: !6)
!5 = !DIFile(filename: "/Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/i386/_types.h", directory: "")
!6 = !DIBasicType(name: "unsigned short", size: 16, encoding: DW_ATE_unsigned)
!7 = !DIDerivedType(tag: DW_TAG_pointer_type, baseType: !8, size: 64)
!8 = distinct !DICompositeType(tag: DW_TAG_structure_type, name: "sockaddr", file: !9, line: 407, size: 128, elements: !10)
!9 = !DIFile(filename: "/Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/sys/socket.h", directory: "")
!10 = !{!11, !14, !17}
!11 = !DIDerivedType(tag: DW_TAG_member, name: "sa_len", scope: !8, file: !9, line: 408, baseType: !12, size: 8)
!12 = !DIDerivedType(tag: DW_TAG_typedef, name: "__uint8_t", file: !5, line: 41, baseType: !13)
!13 = !DIBasicType(name: "unsigned char", size: 8, encoding: DW_ATE_unsigned_char)
!14 = !DIDerivedType(tag: DW_TAG_member, name: "sa_family", scope: !8, file: !9, line: 409, baseType: !15, size: 8, offset: 8)
!15 = !DIDerivedType(tag: DW_TAG_typedef, name: "sa_family_t", file: !16, line: 31, baseType: !12)
!16 = !DIFile(filename: "/Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/sys/_types/_sa_family_t.h", directory: "")
!17 = !DIDerivedType(tag: DW_TAG_member, name: "sa_data", scope: !8, file: !9, line: 410, baseType: !18, size: 112, offset: 16)
!18 = !DICompositeType(tag: DW_TAG_array_type, baseType: !19, size: 112, elements: !20)
!19 = !DIBasicType(name: "char", size: 8, encoding: DW_ATE_signed_char)
!20 = !{!21}
!21 = !DISubrange(count: 14)
!22 = !DIBasicType(name: "unsigned int", size: 32, encoding: DW_ATE_unsigned)
!23 = !{i32 7, !"Dwarf Version", i32 4}
!24 = !{i32 2, !"Debug Info Version", i32 3}
!25 = !{i32 1, !"wchar_size", i32 4}
!26 = !{i32 7, !"PIC Level", i32 2}
!27 = !{!"clang version 10.0.0 "}
!28 = distinct !DISubprogram(name: "CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_11_bad", scope: !1, file: !1, line: 44, type: !29, scopeLine: 45, spFlags: DISPFlagDefinition, unit: !0, retainedNodes: !2)
!29 = !DISubroutineType(types: !30)
!30 = !{null}
!31 = !DILocalVariable(name: "data", scope: !28, file: !1, line: 46, type: !32)
!32 = !DIBasicType(name: "int", size: 32, encoding: DW_ATE_signed)
!33 = !DILocation(line: 46, column: 9, scope: !28)
!34 = !DILocation(line: 48, column: 10, scope: !28)
!35 = !DILocation(line: 49, column: 8, scope: !36)
!36 = distinct !DILexicalBlock(scope: !28, file: !1, line: 49, column: 8)
!37 = !DILocation(line: 49, column: 8, scope: !28)
!38 = !DILocalVariable(name: "recvResult", scope: !39, file: !1, line: 56, type: !32)
!39 = distinct !DILexicalBlock(scope: !40, file: !1, line: 51, column: 9)
!40 = distinct !DILexicalBlock(scope: !36, file: !1, line: 50, column: 5)
!41 = !DILocation(line: 56, column: 17, scope: !39)
!42 = !DILocalVariable(name: "service", scope: !39, file: !1, line: 57, type: !43)
!43 = distinct !DICompositeType(tag: DW_TAG_structure_type, name: "sockaddr_in", file: !44, line: 375, size: 128, elements: !45)
!44 = !DIFile(filename: "/Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/netinet/in.h", directory: "")
!45 = !{!46, !47, !48, !51, !58}
!46 = !DIDerivedType(tag: DW_TAG_member, name: "sin_len", scope: !43, file: !44, line: 376, baseType: !12, size: 8)
!47 = !DIDerivedType(tag: DW_TAG_member, name: "sin_family", scope: !43, file: !44, line: 377, baseType: !15, size: 8, offset: 8)
!48 = !DIDerivedType(tag: DW_TAG_member, name: "sin_port", scope: !43, file: !44, line: 378, baseType: !49, size: 16, offset: 16)
!49 = !DIDerivedType(tag: DW_TAG_typedef, name: "in_port_t", file: !50, line: 31, baseType: !4)
!50 = !DIFile(filename: "/Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/sys/_types/_in_port_t.h", directory: "")
!51 = !DIDerivedType(tag: DW_TAG_member, name: "sin_addr", scope: !43, file: !44, line: 379, baseType: !52, size: 32, offset: 32)
!52 = distinct !DICompositeType(tag: DW_TAG_structure_type, name: "in_addr", file: !44, line: 301, size: 32, elements: !53)
!53 = !{!54}
!54 = !DIDerivedType(tag: DW_TAG_member, name: "s_addr", scope: !52, file: !44, line: 302, baseType: !55, size: 32)
!55 = !DIDerivedType(tag: DW_TAG_typedef, name: "in_addr_t", file: !56, line: 31, baseType: !57)
!56 = !DIFile(filename: "/Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/sys/_types/_in_addr_t.h", directory: "")
!57 = !DIDerivedType(tag: DW_TAG_typedef, name: "__uint32_t", file: !5, line: 45, baseType: !22)
!58 = !DIDerivedType(tag: DW_TAG_member, name: "sin_zero", scope: !43, file: !44, line: 380, baseType: !59, size: 64, offset: 64)
!59 = !DICompositeType(tag: DW_TAG_array_type, baseType: !19, size: 64, elements: !60)
!60 = !{!61}
!61 = !DISubrange(count: 8)
!62 = !DILocation(line: 57, column: 32, scope: !39)
!63 = !DILocalVariable(name: "connectSocket", scope: !39, file: !1, line: 58, type: !32)
!64 = !DILocation(line: 58, column: 20, scope: !39)
!65 = !DILocalVariable(name: "inputBuffer", scope: !39, file: !1, line: 59, type: !18)
!66 = !DILocation(line: 59, column: 18, scope: !39)
!67 = !DILocation(line: 60, column: 13, scope: !39)
!68 = !DILocation(line: 70, column: 33, scope: !69)
!69 = distinct !DILexicalBlock(scope: !39, file: !1, line: 61, column: 13)
!70 = !DILocation(line: 70, column: 31, scope: !69)
!71 = !DILocation(line: 71, column: 21, scope: !72)
!72 = distinct !DILexicalBlock(scope: !69, file: !1, line: 71, column: 21)
!73 = !DILocation(line: 71, column: 35, scope: !72)
!74 = !DILocation(line: 71, column: 21, scope: !69)
!75 = !DILocation(line: 73, column: 21, scope: !76)
!76 = distinct !DILexicalBlock(scope: !72, file: !1, line: 72, column: 17)
!77 = !DILocation(line: 75, column: 17, scope: !69)
!78 = !DILocation(line: 76, column: 25, scope: !69)
!79 = !DILocation(line: 76, column: 36, scope: !69)
!80 = !DILocation(line: 77, column: 43, scope: !69)
!81 = !DILocation(line: 77, column: 25, scope: !69)
!82 = !DILocation(line: 77, column: 34, scope: !69)
!83 = !DILocation(line: 77, column: 41, scope: !69)
!84 = !DILocation(line: 78, column: 25, scope: !69)
!85 = !DILocation(line: 78, column: 34, scope: !69)
!86 = !DILocation(line: 79, column: 29, scope: !87)
!87 = distinct !DILexicalBlock(scope: !69, file: !1, line: 79, column: 21)
!88 = !DILocation(line: 79, column: 44, scope: !87)
!89 = !DILocation(line: 79, column: 21, scope: !87)
!90 = !DILocation(line: 79, column: 89, scope: !87)
!91 = !DILocation(line: 79, column: 21, scope: !69)
!92 = !DILocation(line: 81, column: 21, scope: !93)
!93 = distinct !DILexicalBlock(scope: !87, file: !1, line: 80, column: 17)
!94 = !DILocation(line: 85, column: 35, scope: !69)
!95 = !DILocation(line: 85, column: 50, scope: !69)
!96 = !DILocation(line: 85, column: 30, scope: !69)
!97 = !DILocation(line: 85, column: 28, scope: !69)
!98 = !DILocation(line: 86, column: 21, scope: !99)
!99 = distinct !DILexicalBlock(scope: !69, file: !1, line: 86, column: 21)
!100 = !DILocation(line: 86, column: 32, scope: !99)
!101 = !DILocation(line: 86, column: 48, scope: !99)
!102 = !DILocation(line: 86, column: 51, scope: !99)
!103 = !DILocation(line: 86, column: 62, scope: !99)
!104 = !DILocation(line: 86, column: 21, scope: !69)
!105 = !DILocation(line: 88, column: 21, scope: !106)
!106 = distinct !DILexicalBlock(scope: !99, file: !1, line: 87, column: 17)
!107 = !DILocation(line: 91, column: 29, scope: !69)
!108 = !DILocation(line: 91, column: 17, scope: !69)
!109 = !DILocation(line: 91, column: 41, scope: !69)
!110 = !DILocation(line: 93, column: 29, scope: !69)
!111 = !DILocation(line: 93, column: 24, scope: !69)
!112 = !DILocation(line: 93, column: 22, scope: !69)
!113 = !DILocation(line: 94, column: 13, scope: !69)
!114 = !DILocation(line: 96, column: 17, scope: !115)
!115 = distinct !DILexicalBlock(scope: !39, file: !1, line: 96, column: 17)
!116 = !DILocation(line: 96, column: 31, scope: !115)
!117 = !DILocation(line: 96, column: 17, scope: !39)
!118 = !DILocation(line: 98, column: 30, scope: !119)
!119 = distinct !DILexicalBlock(scope: !115, file: !1, line: 97, column: 13)
!120 = !DILocation(line: 98, column: 17, scope: !119)
!121 = !DILocation(line: 99, column: 13, scope: !119)
!122 = !DILocation(line: 107, column: 5, scope: !40)
!123 = !DILocation(line: 108, column: 8, scope: !124)
!124 = distinct !DILexicalBlock(scope: !28, file: !1, line: 108, column: 8)
!125 = !DILocation(line: 108, column: 8, scope: !28)
!126 = !DILocalVariable(name: "i", scope: !127, file: !1, line: 111, type: !32)
!127 = distinct !DILexicalBlock(scope: !128, file: !1, line: 110, column: 9)
!128 = distinct !DILexicalBlock(scope: !124, file: !1, line: 109, column: 5)
!129 = !DILocation(line: 111, column: 17, scope: !127)
!130 = !DILocalVariable(name: "buffer", scope: !127, file: !1, line: 112, type: !131)
!131 = !DICompositeType(tag: DW_TAG_array_type, baseType: !32, size: 320, elements: !132)
!132 = !{!133}
!133 = !DISubrange(count: 10)
!134 = !DILocation(line: 112, column: 17, scope: !127)
!135 = !DILocation(line: 115, column: 17, scope: !136)
!136 = distinct !DILexicalBlock(scope: !127, file: !1, line: 115, column: 17)
!137 = !DILocation(line: 115, column: 22, scope: !136)
!138 = !DILocation(line: 115, column: 17, scope: !127)
!139 = !DILocation(line: 117, column: 24, scope: !140)
!140 = distinct !DILexicalBlock(scope: !136, file: !1, line: 116, column: 13)
!141 = !DILocation(line: 117, column: 17, scope: !140)
!142 = !DILocation(line: 117, column: 30, scope: !140)
!143 = !DILocation(line: 119, column: 23, scope: !144)
!144 = distinct !DILexicalBlock(scope: !140, file: !1, line: 119, column: 17)
!145 = !DILocation(line: 119, column: 21, scope: !144)
!146 = !DILocation(line: 119, column: 28, scope: !147)
!147 = distinct !DILexicalBlock(scope: !144, file: !1, line: 119, column: 17)
!148 = !DILocation(line: 119, column: 30, scope: !147)
!149 = !DILocation(line: 119, column: 17, scope: !144)
!150 = !DILocation(line: 121, column: 41, scope: !151)
!151 = distinct !DILexicalBlock(scope: !147, file: !1, line: 120, column: 17)
!152 = !DILocation(line: 121, column: 34, scope: !151)
!153 = !DILocation(line: 121, column: 21, scope: !151)
!154 = !DILocation(line: 122, column: 17, scope: !151)
!155 = !DILocation(line: 119, column: 37, scope: !147)
!156 = !DILocation(line: 119, column: 17, scope: !147)
!157 = distinct !{!157, !149, !158}
!158 = !DILocation(line: 122, column: 17, scope: !144)
!159 = !DILocation(line: 123, column: 13, scope: !140)
!160 = !DILocation(line: 126, column: 17, scope: !161)
!161 = distinct !DILexicalBlock(scope: !136, file: !1, line: 125, column: 13)
!162 = !DILocation(line: 129, column: 5, scope: !128)
!163 = !DILocation(line: 130, column: 1, scope: !28)
!164 = distinct !DISubprogram(name: "CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_11_good", scope: !1, file: !1, line: 394, type: !29, scopeLine: 395, spFlags: DISPFlagDefinition, unit: !0, retainedNodes: !2)
!165 = !DILocation(line: 396, column: 5, scope: !164)
!166 = !DILocation(line: 397, column: 5, scope: !164)
!167 = !DILocation(line: 398, column: 5, scope: !164)
!168 = !DILocation(line: 399, column: 5, scope: !164)
!169 = !DILocation(line: 400, column: 1, scope: !164)
!170 = distinct !DISubprogram(name: "main", scope: !1, file: !1, line: 411, type: !171, scopeLine: 412, flags: DIFlagPrototyped, spFlags: DISPFlagDefinition, unit: !0, retainedNodes: !2)
!171 = !DISubroutineType(types: !172)
!172 = !{!32, !32, !173}
!173 = !DIDerivedType(tag: DW_TAG_pointer_type, baseType: !174, size: 64)
!174 = !DIDerivedType(tag: DW_TAG_pointer_type, baseType: !19, size: 64)
!175 = !DILocalVariable(name: "argc", arg: 1, scope: !170, file: !1, line: 411, type: !32)
!176 = !DILocation(line: 411, column: 14, scope: !170)
!177 = !DILocalVariable(name: "argv", arg: 2, scope: !170, file: !1, line: 411, type: !173)
!178 = !DILocation(line: 411, column: 27, scope: !170)
!179 = !DILocation(line: 414, column: 22, scope: !170)
!180 = !DILocation(line: 414, column: 12, scope: !170)
!181 = !DILocation(line: 414, column: 5, scope: !170)
!182 = !DILocation(line: 416, column: 5, scope: !170)
!183 = !DILocation(line: 417, column: 5, scope: !170)
!184 = !DILocation(line: 418, column: 5, scope: !170)
!185 = !DILocation(line: 421, column: 5, scope: !170)
!186 = !DILocation(line: 422, column: 5, scope: !170)
!187 = !DILocation(line: 423, column: 5, scope: !170)
!188 = !DILocation(line: 425, column: 5, scope: !170)
!189 = distinct !DISubprogram(name: "goodB2G1", scope: !1, file: !1, line: 137, type: !29, scopeLine: 138, spFlags: DISPFlagLocalToUnit | DISPFlagDefinition, unit: !0, retainedNodes: !2)
!190 = !DILocalVariable(name: "data", scope: !189, file: !1, line: 139, type: !32)
!191 = !DILocation(line: 139, column: 9, scope: !189)
!192 = !DILocation(line: 141, column: 10, scope: !189)
!193 = !DILocation(line: 142, column: 8, scope: !194)
!194 = distinct !DILexicalBlock(scope: !189, file: !1, line: 142, column: 8)
!195 = !DILocation(line: 142, column: 8, scope: !189)
!196 = !DILocalVariable(name: "recvResult", scope: !197, file: !1, line: 149, type: !32)
!197 = distinct !DILexicalBlock(scope: !198, file: !1, line: 144, column: 9)
!198 = distinct !DILexicalBlock(scope: !194, file: !1, line: 143, column: 5)
!199 = !DILocation(line: 149, column: 17, scope: !197)
!200 = !DILocalVariable(name: "service", scope: !197, file: !1, line: 150, type: !43)
!201 = !DILocation(line: 150, column: 32, scope: !197)
!202 = !DILocalVariable(name: "connectSocket", scope: !197, file: !1, line: 151, type: !32)
!203 = !DILocation(line: 151, column: 20, scope: !197)
!204 = !DILocalVariable(name: "inputBuffer", scope: !197, file: !1, line: 152, type: !18)
!205 = !DILocation(line: 152, column: 18, scope: !197)
!206 = !DILocation(line: 153, column: 13, scope: !197)
!207 = !DILocation(line: 163, column: 33, scope: !208)
!208 = distinct !DILexicalBlock(scope: !197, file: !1, line: 154, column: 13)
!209 = !DILocation(line: 163, column: 31, scope: !208)
!210 = !DILocation(line: 164, column: 21, scope: !211)
!211 = distinct !DILexicalBlock(scope: !208, file: !1, line: 164, column: 21)
!212 = !DILocation(line: 164, column: 35, scope: !211)
!213 = !DILocation(line: 164, column: 21, scope: !208)
!214 = !DILocation(line: 166, column: 21, scope: !215)
!215 = distinct !DILexicalBlock(scope: !211, file: !1, line: 165, column: 17)
!216 = !DILocation(line: 168, column: 17, scope: !208)
!217 = !DILocation(line: 169, column: 25, scope: !208)
!218 = !DILocation(line: 169, column: 36, scope: !208)
!219 = !DILocation(line: 170, column: 43, scope: !208)
!220 = !DILocation(line: 170, column: 25, scope: !208)
!221 = !DILocation(line: 170, column: 34, scope: !208)
!222 = !DILocation(line: 170, column: 41, scope: !208)
!223 = !DILocation(line: 171, column: 25, scope: !208)
!224 = !DILocation(line: 171, column: 34, scope: !208)
!225 = !DILocation(line: 172, column: 29, scope: !226)
!226 = distinct !DILexicalBlock(scope: !208, file: !1, line: 172, column: 21)
!227 = !DILocation(line: 172, column: 44, scope: !226)
!228 = !DILocation(line: 172, column: 21, scope: !226)
!229 = !DILocation(line: 172, column: 89, scope: !226)
!230 = !DILocation(line: 172, column: 21, scope: !208)
!231 = !DILocation(line: 174, column: 21, scope: !232)
!232 = distinct !DILexicalBlock(scope: !226, file: !1, line: 173, column: 17)
!233 = !DILocation(line: 178, column: 35, scope: !208)
!234 = !DILocation(line: 178, column: 50, scope: !208)
!235 = !DILocation(line: 178, column: 30, scope: !208)
!236 = !DILocation(line: 178, column: 28, scope: !208)
!237 = !DILocation(line: 179, column: 21, scope: !238)
!238 = distinct !DILexicalBlock(scope: !208, file: !1, line: 179, column: 21)
!239 = !DILocation(line: 179, column: 32, scope: !238)
!240 = !DILocation(line: 179, column: 48, scope: !238)
!241 = !DILocation(line: 179, column: 51, scope: !238)
!242 = !DILocation(line: 179, column: 62, scope: !238)
!243 = !DILocation(line: 179, column: 21, scope: !208)
!244 = !DILocation(line: 181, column: 21, scope: !245)
!245 = distinct !DILexicalBlock(scope: !238, file: !1, line: 180, column: 17)
!246 = !DILocation(line: 184, column: 29, scope: !208)
!247 = !DILocation(line: 184, column: 17, scope: !208)
!248 = !DILocation(line: 184, column: 41, scope: !208)
!249 = !DILocation(line: 186, column: 29, scope: !208)
!250 = !DILocation(line: 186, column: 24, scope: !208)
!251 = !DILocation(line: 186, column: 22, scope: !208)
!252 = !DILocation(line: 187, column: 13, scope: !208)
!253 = !DILocation(line: 189, column: 17, scope: !254)
!254 = distinct !DILexicalBlock(scope: !197, file: !1, line: 189, column: 17)
!255 = !DILocation(line: 189, column: 31, scope: !254)
!256 = !DILocation(line: 189, column: 17, scope: !197)
!257 = !DILocation(line: 191, column: 30, scope: !258)
!258 = distinct !DILexicalBlock(scope: !254, file: !1, line: 190, column: 13)
!259 = !DILocation(line: 191, column: 17, scope: !258)
!260 = !DILocation(line: 192, column: 13, scope: !258)
!261 = !DILocation(line: 200, column: 5, scope: !198)
!262 = !DILocation(line: 201, column: 8, scope: !263)
!263 = distinct !DILexicalBlock(scope: !189, file: !1, line: 201, column: 8)
!264 = !DILocation(line: 201, column: 8, scope: !189)
!265 = !DILocation(line: 204, column: 9, scope: !266)
!266 = distinct !DILexicalBlock(scope: !263, file: !1, line: 202, column: 5)
!267 = !DILocation(line: 205, column: 5, scope: !266)
!268 = !DILocalVariable(name: "i", scope: !269, file: !1, line: 209, type: !32)
!269 = distinct !DILexicalBlock(scope: !270, file: !1, line: 208, column: 9)
!270 = distinct !DILexicalBlock(scope: !263, file: !1, line: 207, column: 5)
!271 = !DILocation(line: 209, column: 17, scope: !269)
!272 = !DILocalVariable(name: "buffer", scope: !269, file: !1, line: 210, type: !131)
!273 = !DILocation(line: 210, column: 17, scope: !269)
!274 = !DILocation(line: 212, column: 17, scope: !275)
!275 = distinct !DILexicalBlock(scope: !269, file: !1, line: 212, column: 17)
!276 = !DILocation(line: 212, column: 22, scope: !275)
!277 = !DILocation(line: 212, column: 27, scope: !275)
!278 = !DILocation(line: 212, column: 30, scope: !275)
!279 = !DILocation(line: 212, column: 35, scope: !275)
!280 = !DILocation(line: 212, column: 17, scope: !269)
!281 = !DILocation(line: 214, column: 24, scope: !282)
!282 = distinct !DILexicalBlock(scope: !275, file: !1, line: 213, column: 13)
!283 = !DILocation(line: 214, column: 17, scope: !282)
!284 = !DILocation(line: 214, column: 30, scope: !282)
!285 = !DILocation(line: 216, column: 23, scope: !286)
!286 = distinct !DILexicalBlock(scope: !282, file: !1, line: 216, column: 17)
!287 = !DILocation(line: 216, column: 21, scope: !286)
!288 = !DILocation(line: 216, column: 28, scope: !289)
!289 = distinct !DILexicalBlock(scope: !286, file: !1, line: 216, column: 17)
!290 = !DILocation(line: 216, column: 30, scope: !289)
!291 = !DILocation(line: 216, column: 17, scope: !286)
!292 = !DILocation(line: 218, column: 41, scope: !293)
!293 = distinct !DILexicalBlock(scope: !289, file: !1, line: 217, column: 17)
!294 = !DILocation(line: 218, column: 34, scope: !293)
!295 = !DILocation(line: 218, column: 21, scope: !293)
!296 = !DILocation(line: 219, column: 17, scope: !293)
!297 = !DILocation(line: 216, column: 37, scope: !289)
!298 = !DILocation(line: 216, column: 17, scope: !289)
!299 = distinct !{!299, !291, !300}
!300 = !DILocation(line: 219, column: 17, scope: !286)
!301 = !DILocation(line: 220, column: 13, scope: !282)
!302 = !DILocation(line: 223, column: 17, scope: !303)
!303 = distinct !DILexicalBlock(scope: !275, file: !1, line: 222, column: 13)
!304 = !DILocation(line: 227, column: 1, scope: !189)
!305 = distinct !DISubprogram(name: "goodB2G2", scope: !1, file: !1, line: 230, type: !29, scopeLine: 231, spFlags: DISPFlagLocalToUnit | DISPFlagDefinition, unit: !0, retainedNodes: !2)
!306 = !DILocalVariable(name: "data", scope: !305, file: !1, line: 232, type: !32)
!307 = !DILocation(line: 232, column: 9, scope: !305)
!308 = !DILocation(line: 234, column: 10, scope: !305)
!309 = !DILocation(line: 235, column: 8, scope: !310)
!310 = distinct !DILexicalBlock(scope: !305, file: !1, line: 235, column: 8)
!311 = !DILocation(line: 235, column: 8, scope: !305)
!312 = !DILocalVariable(name: "recvResult", scope: !313, file: !1, line: 242, type: !32)
!313 = distinct !DILexicalBlock(scope: !314, file: !1, line: 237, column: 9)
!314 = distinct !DILexicalBlock(scope: !310, file: !1, line: 236, column: 5)
!315 = !DILocation(line: 242, column: 17, scope: !313)
!316 = !DILocalVariable(name: "service", scope: !313, file: !1, line: 243, type: !43)
!317 = !DILocation(line: 243, column: 32, scope: !313)
!318 = !DILocalVariable(name: "connectSocket", scope: !313, file: !1, line: 244, type: !32)
!319 = !DILocation(line: 244, column: 20, scope: !313)
!320 = !DILocalVariable(name: "inputBuffer", scope: !313, file: !1, line: 245, type: !18)
!321 = !DILocation(line: 245, column: 18, scope: !313)
!322 = !DILocation(line: 246, column: 13, scope: !313)
!323 = !DILocation(line: 256, column: 33, scope: !324)
!324 = distinct !DILexicalBlock(scope: !313, file: !1, line: 247, column: 13)
!325 = !DILocation(line: 256, column: 31, scope: !324)
!326 = !DILocation(line: 257, column: 21, scope: !327)
!327 = distinct !DILexicalBlock(scope: !324, file: !1, line: 257, column: 21)
!328 = !DILocation(line: 257, column: 35, scope: !327)
!329 = !DILocation(line: 257, column: 21, scope: !324)
!330 = !DILocation(line: 259, column: 21, scope: !331)
!331 = distinct !DILexicalBlock(scope: !327, file: !1, line: 258, column: 17)
!332 = !DILocation(line: 261, column: 17, scope: !324)
!333 = !DILocation(line: 262, column: 25, scope: !324)
!334 = !DILocation(line: 262, column: 36, scope: !324)
!335 = !DILocation(line: 263, column: 43, scope: !324)
!336 = !DILocation(line: 263, column: 25, scope: !324)
!337 = !DILocation(line: 263, column: 34, scope: !324)
!338 = !DILocation(line: 263, column: 41, scope: !324)
!339 = !DILocation(line: 264, column: 25, scope: !324)
!340 = !DILocation(line: 264, column: 34, scope: !324)
!341 = !DILocation(line: 265, column: 29, scope: !342)
!342 = distinct !DILexicalBlock(scope: !324, file: !1, line: 265, column: 21)
!343 = !DILocation(line: 265, column: 44, scope: !342)
!344 = !DILocation(line: 265, column: 21, scope: !342)
!345 = !DILocation(line: 265, column: 89, scope: !342)
!346 = !DILocation(line: 265, column: 21, scope: !324)
!347 = !DILocation(line: 267, column: 21, scope: !348)
!348 = distinct !DILexicalBlock(scope: !342, file: !1, line: 266, column: 17)
!349 = !DILocation(line: 271, column: 35, scope: !324)
!350 = !DILocation(line: 271, column: 50, scope: !324)
!351 = !DILocation(line: 271, column: 30, scope: !324)
!352 = !DILocation(line: 271, column: 28, scope: !324)
!353 = !DILocation(line: 272, column: 21, scope: !354)
!354 = distinct !DILexicalBlock(scope: !324, file: !1, line: 272, column: 21)
!355 = !DILocation(line: 272, column: 32, scope: !354)
!356 = !DILocation(line: 272, column: 48, scope: !354)
!357 = !DILocation(line: 272, column: 51, scope: !354)
!358 = !DILocation(line: 272, column: 62, scope: !354)
!359 = !DILocation(line: 272, column: 21, scope: !324)
!360 = !DILocation(line: 274, column: 21, scope: !361)
!361 = distinct !DILexicalBlock(scope: !354, file: !1, line: 273, column: 17)
!362 = !DILocation(line: 277, column: 29, scope: !324)
!363 = !DILocation(line: 277, column: 17, scope: !324)
!364 = !DILocation(line: 277, column: 41, scope: !324)
!365 = !DILocation(line: 279, column: 29, scope: !324)
!366 = !DILocation(line: 279, column: 24, scope: !324)
!367 = !DILocation(line: 279, column: 22, scope: !324)
!368 = !DILocation(line: 280, column: 13, scope: !324)
!369 = !DILocation(line: 282, column: 17, scope: !370)
!370 = distinct !DILexicalBlock(scope: !313, file: !1, line: 282, column: 17)
!371 = !DILocation(line: 282, column: 31, scope: !370)
!372 = !DILocation(line: 282, column: 17, scope: !313)
!373 = !DILocation(line: 284, column: 30, scope: !374)
!374 = distinct !DILexicalBlock(scope: !370, file: !1, line: 283, column: 13)
!375 = !DILocation(line: 284, column: 17, scope: !374)
!376 = !DILocation(line: 285, column: 13, scope: !374)
!377 = !DILocation(line: 293, column: 5, scope: !314)
!378 = !DILocation(line: 294, column: 8, scope: !379)
!379 = distinct !DILexicalBlock(scope: !305, file: !1, line: 294, column: 8)
!380 = !DILocation(line: 294, column: 8, scope: !305)
!381 = !DILocalVariable(name: "i", scope: !382, file: !1, line: 297, type: !32)
!382 = distinct !DILexicalBlock(scope: !383, file: !1, line: 296, column: 9)
!383 = distinct !DILexicalBlock(scope: !379, file: !1, line: 295, column: 5)
!384 = !DILocation(line: 297, column: 17, scope: !382)
!385 = !DILocalVariable(name: "buffer", scope: !382, file: !1, line: 298, type: !131)
!386 = !DILocation(line: 298, column: 17, scope: !382)
!387 = !DILocation(line: 300, column: 17, scope: !388)
!388 = distinct !DILexicalBlock(scope: !382, file: !1, line: 300, column: 17)
!389 = !DILocation(line: 300, column: 22, scope: !388)
!390 = !DILocation(line: 300, column: 27, scope: !388)
!391 = !DILocation(line: 300, column: 30, scope: !388)
!392 = !DILocation(line: 300, column: 35, scope: !388)
!393 = !DILocation(line: 300, column: 17, scope: !382)
!394 = !DILocation(line: 302, column: 24, scope: !395)
!395 = distinct !DILexicalBlock(scope: !388, file: !1, line: 301, column: 13)
!396 = !DILocation(line: 302, column: 17, scope: !395)
!397 = !DILocation(line: 302, column: 30, scope: !395)
!398 = !DILocation(line: 304, column: 23, scope: !399)
!399 = distinct !DILexicalBlock(scope: !395, file: !1, line: 304, column: 17)
!400 = !DILocation(line: 304, column: 21, scope: !399)
!401 = !DILocation(line: 304, column: 28, scope: !402)
!402 = distinct !DILexicalBlock(scope: !399, file: !1, line: 304, column: 17)
!403 = !DILocation(line: 304, column: 30, scope: !402)
!404 = !DILocation(line: 304, column: 17, scope: !399)
!405 = !DILocation(line: 306, column: 41, scope: !406)
!406 = distinct !DILexicalBlock(scope: !402, file: !1, line: 305, column: 17)
!407 = !DILocation(line: 306, column: 34, scope: !406)
!408 = !DILocation(line: 306, column: 21, scope: !406)
!409 = !DILocation(line: 307, column: 17, scope: !406)
!410 = !DILocation(line: 304, column: 37, scope: !402)
!411 = !DILocation(line: 304, column: 17, scope: !402)
!412 = distinct !{!412, !404, !413}
!413 = !DILocation(line: 307, column: 17, scope: !399)
!414 = !DILocation(line: 308, column: 13, scope: !395)
!415 = !DILocation(line: 311, column: 17, scope: !416)
!416 = distinct !DILexicalBlock(scope: !388, file: !1, line: 310, column: 13)
!417 = !DILocation(line: 314, column: 5, scope: !383)
!418 = !DILocation(line: 315, column: 1, scope: !305)
!419 = distinct !DISubprogram(name: "goodG2B1", scope: !1, file: !1, line: 318, type: !29, scopeLine: 319, spFlags: DISPFlagLocalToUnit | DISPFlagDefinition, unit: !0, retainedNodes: !2)
!420 = !DILocalVariable(name: "data", scope: !419, file: !1, line: 320, type: !32)
!421 = !DILocation(line: 320, column: 9, scope: !419)
!422 = !DILocation(line: 322, column: 10, scope: !419)
!423 = !DILocation(line: 323, column: 8, scope: !424)
!424 = distinct !DILexicalBlock(scope: !419, file: !1, line: 323, column: 8)
!425 = !DILocation(line: 323, column: 8, scope: !419)
!426 = !DILocation(line: 326, column: 9, scope: !427)
!427 = distinct !DILexicalBlock(scope: !424, file: !1, line: 324, column: 5)
!428 = !DILocation(line: 327, column: 5, scope: !427)
!429 = !DILocation(line: 332, column: 14, scope: !430)
!430 = distinct !DILexicalBlock(scope: !424, file: !1, line: 329, column: 5)
!431 = !DILocation(line: 334, column: 8, scope: !432)
!432 = distinct !DILexicalBlock(scope: !419, file: !1, line: 334, column: 8)
!433 = !DILocation(line: 334, column: 8, scope: !419)
!434 = !DILocalVariable(name: "i", scope: !435, file: !1, line: 337, type: !32)
!435 = distinct !DILexicalBlock(scope: !436, file: !1, line: 336, column: 9)
!436 = distinct !DILexicalBlock(scope: !432, file: !1, line: 335, column: 5)
!437 = !DILocation(line: 337, column: 17, scope: !435)
!438 = !DILocalVariable(name: "buffer", scope: !435, file: !1, line: 338, type: !131)
!439 = !DILocation(line: 338, column: 17, scope: !435)
!440 = !DILocation(line: 341, column: 17, scope: !441)
!441 = distinct !DILexicalBlock(scope: !435, file: !1, line: 341, column: 17)
!442 = !DILocation(line: 341, column: 22, scope: !441)
!443 = !DILocation(line: 341, column: 17, scope: !435)
!444 = !DILocation(line: 343, column: 24, scope: !445)
!445 = distinct !DILexicalBlock(scope: !441, file: !1, line: 342, column: 13)
!446 = !DILocation(line: 343, column: 17, scope: !445)
!447 = !DILocation(line: 343, column: 30, scope: !445)
!448 = !DILocation(line: 345, column: 23, scope: !449)
!449 = distinct !DILexicalBlock(scope: !445, file: !1, line: 345, column: 17)
!450 = !DILocation(line: 345, column: 21, scope: !449)
!451 = !DILocation(line: 345, column: 28, scope: !452)
!452 = distinct !DILexicalBlock(scope: !449, file: !1, line: 345, column: 17)
!453 = !DILocation(line: 345, column: 30, scope: !452)
!454 = !DILocation(line: 345, column: 17, scope: !449)
!455 = !DILocation(line: 347, column: 41, scope: !456)
!456 = distinct !DILexicalBlock(scope: !452, file: !1, line: 346, column: 17)
!457 = !DILocation(line: 347, column: 34, scope: !456)
!458 = !DILocation(line: 347, column: 21, scope: !456)
!459 = !DILocation(line: 348, column: 17, scope: !456)
!460 = !DILocation(line: 345, column: 37, scope: !452)
!461 = !DILocation(line: 345, column: 17, scope: !452)
!462 = distinct !{!462, !454, !463}
!463 = !DILocation(line: 348, column: 17, scope: !449)
!464 = !DILocation(line: 349, column: 13, scope: !445)
!465 = !DILocation(line: 352, column: 17, scope: !466)
!466 = distinct !DILexicalBlock(scope: !441, file: !1, line: 351, column: 13)
!467 = !DILocation(line: 355, column: 5, scope: !436)
!468 = !DILocation(line: 356, column: 1, scope: !419)
!469 = distinct !DISubprogram(name: "goodG2B2", scope: !1, file: !1, line: 359, type: !29, scopeLine: 360, spFlags: DISPFlagLocalToUnit | DISPFlagDefinition, unit: !0, retainedNodes: !2)
!470 = !DILocalVariable(name: "data", scope: !469, file: !1, line: 361, type: !32)
!471 = !DILocation(line: 361, column: 9, scope: !469)
!472 = !DILocation(line: 363, column: 10, scope: !469)
!473 = !DILocation(line: 364, column: 8, scope: !474)
!474 = distinct !DILexicalBlock(scope: !469, file: !1, line: 364, column: 8)
!475 = !DILocation(line: 364, column: 8, scope: !469)
!476 = !DILocation(line: 368, column: 14, scope: !477)
!477 = distinct !DILexicalBlock(scope: !474, file: !1, line: 365, column: 5)
!478 = !DILocation(line: 369, column: 5, scope: !477)
!479 = !DILocation(line: 370, column: 8, scope: !480)
!480 = distinct !DILexicalBlock(scope: !469, file: !1, line: 370, column: 8)
!481 = !DILocation(line: 370, column: 8, scope: !469)
!482 = !DILocalVariable(name: "i", scope: !483, file: !1, line: 373, type: !32)
!483 = distinct !DILexicalBlock(scope: !484, file: !1, line: 372, column: 9)
!484 = distinct !DILexicalBlock(scope: !480, file: !1, line: 371, column: 5)
!485 = !DILocation(line: 373, column: 17, scope: !483)
!486 = !DILocalVariable(name: "buffer", scope: !483, file: !1, line: 374, type: !131)
!487 = !DILocation(line: 374, column: 17, scope: !483)
!488 = !DILocation(line: 377, column: 17, scope: !489)
!489 = distinct !DILexicalBlock(scope: !483, file: !1, line: 377, column: 17)
!490 = !DILocation(line: 377, column: 22, scope: !489)
!491 = !DILocation(line: 377, column: 17, scope: !483)
!492 = !DILocation(line: 379, column: 24, scope: !493)
!493 = distinct !DILexicalBlock(scope: !489, file: !1, line: 378, column: 13)
!494 = !DILocation(line: 379, column: 17, scope: !493)
!495 = !DILocation(line: 379, column: 30, scope: !493)
!496 = !DILocation(line: 381, column: 23, scope: !497)
!497 = distinct !DILexicalBlock(scope: !493, file: !1, line: 381, column: 17)
!498 = !DILocation(line: 381, column: 21, scope: !497)
!499 = !DILocation(line: 381, column: 28, scope: !500)
!500 = distinct !DILexicalBlock(scope: !497, file: !1, line: 381, column: 17)
!501 = !DILocation(line: 381, column: 30, scope: !500)
!502 = !DILocation(line: 381, column: 17, scope: !497)
!503 = !DILocation(line: 383, column: 41, scope: !504)
!504 = distinct !DILexicalBlock(scope: !500, file: !1, line: 382, column: 17)
!505 = !DILocation(line: 383, column: 34, scope: !504)
!506 = !DILocation(line: 383, column: 21, scope: !504)
!507 = !DILocation(line: 384, column: 17, scope: !504)
!508 = !DILocation(line: 381, column: 37, scope: !500)
!509 = !DILocation(line: 381, column: 17, scope: !500)
!510 = distinct !{!510, !502, !511}
!511 = !DILocation(line: 384, column: 17, scope: !497)
!512 = !DILocation(line: 385, column: 13, scope: !493)
!513 = !DILocation(line: 388, column: 17, scope: !514)
!514 = distinct !DILexicalBlock(scope: !489, file: !1, line: 387, column: 13)
!515 = !DILocation(line: 391, column: 5, scope: !484)
!516 = !DILocation(line: 392, column: 1, scope: !469)

