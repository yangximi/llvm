; ModuleID = 'CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c'
target datalayout = "e-p:64:64:64-i1:8:8-i8:8:8-i16:16:16-i32:32:32-i64:64:64-f32:32:32-f64:64:64-v64:64:64-v128:128:128-a0:0:64-s0:64:64-f80:128:128-n8:16:32:64-S128"
target triple = "x86_64-redhat-linux-gnu"

%struct.sockaddr_in = type { i16, i16, %struct.in_addr, [8 x i8] }
%struct.in_addr = type { i32 }
%struct.sockaddr = type { i16, [14 x i8] }

@.str = private unnamed_addr constant [10 x i8] c"127.0.0.1\00", align 1
@.str1 = private unnamed_addr constant [32 x i8] c"ERROR: Array index is negative.\00", align 1
@.str2 = private unnamed_addr constant [36 x i8] c"ERROR: Array index is out-of-bounds\00", align 1

; Function Attrs: nounwind uwtable
define void @CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_bad() #0 {
  %data = alloca i32, align 4
  %recvResult = alloca i32, align 4
  %service = alloca %struct.sockaddr_in, align 4
  %connectSocket = alloca i32, align 4
  %inputBuffer = alloca [14 x i8], align 1
  %i = alloca i32, align 4
  %buffer = alloca [10 x i32], align 16
  call void @llvm.dbg.declare(metadata !{i32* %data}, metadata !55), !dbg !57
  store i32 -1, i32* %data, align 4, !dbg !58
  call void @llvm.dbg.declare(metadata !{i32* %recvResult}, metadata !59), !dbg !61
  call void @llvm.dbg.declare(metadata !{%struct.sockaddr_in* %service}, metadata !62), !dbg !83
  call void @llvm.dbg.declare(metadata !{i32* %connectSocket}, metadata !84), !dbg !85
  store i32 -1, i32* %connectSocket, align 4, !dbg !85
  call void @llvm.dbg.declare(metadata !{[14 x i8]* %inputBuffer}, metadata !86), !dbg !91
  br label %1, !dbg !92

; <label>:1                                       ; preds = %0
  %2 = call i32 @socket(i32 2, i32 1, i32 6) #3, !dbg !93
  store i32 %2, i32* %connectSocket, align 4, !dbg !93
  %3 = load i32* %connectSocket, align 4, !dbg !95
  %4 = icmp eq i32 %3, -1, !dbg !95
  br i1 %4, label %5, label %6, !dbg !95

; <label>:5                                       ; preds = %1
  br label %36, !dbg !97

; <label>:6                                       ; preds = %1
  %7 = bitcast %struct.sockaddr_in* %service to i8*, !dbg !99
  call void @llvm.memset.p0i8.i64(i8* %7, i8 0, i64 16, i32 4, i1 false), !dbg !99
  %8 = getelementptr inbounds %struct.sockaddr_in* %service, i32 0, i32 0, !dbg !100
  store i16 2, i16* %8, align 2, !dbg !100
  %9 = call i32 @inet_addr(i8* getelementptr inbounds ([10 x i8]* @.str, i32 0, i32 0)) #3, !dbg !101
  %10 = getelementptr inbounds %struct.sockaddr_in* %service, i32 0, i32 2, !dbg !101
  %11 = getelementptr inbounds %struct.in_addr* %10, i32 0, i32 0, !dbg !101
  store i32 %9, i32* %11, align 4, !dbg !101
  %12 = call zeroext i16 @htons(i16 zeroext 27015) #1, !dbg !102
  %13 = getelementptr inbounds %struct.sockaddr_in* %service, i32 0, i32 1, !dbg !102
  store i16 %12, i16* %13, align 2, !dbg !102
  %14 = load i32* %connectSocket, align 4, !dbg !103
  %15 = bitcast %struct.sockaddr_in* %service to %struct.sockaddr*, !dbg !103
  %16 = call i32 @connect(i32 %14, %struct.sockaddr* %15, i32 16), !dbg !103
  %17 = icmp eq i32 %16, -1, !dbg !103
  br i1 %17, label %18, label %19, !dbg !103

; <label>:18                                      ; preds = %6
  br label %36, !dbg !105

; <label>:19                                      ; preds = %6
  %20 = load i32* %connectSocket, align 4, !dbg !107
  %21 = getelementptr inbounds [14 x i8]* %inputBuffer, i32 0, i32 0, !dbg !107
  %22 = call i64 @recv(i32 %20, i8* %21, i64 13, i32 0), !dbg !107
  %23 = trunc i64 %22 to i32, !dbg !107
  store i32 %23, i32* %recvResult, align 4, !dbg !107
  %24 = load i32* %recvResult, align 4, !dbg !108
  %25 = icmp eq i32 %24, -1, !dbg !108
  br i1 %25, label %29, label %26, !dbg !108

; <label>:26                                      ; preds = %19
  %27 = load i32* %recvResult, align 4, !dbg !108
  %28 = icmp eq i32 %27, 0, !dbg !108
  br i1 %28, label %29, label %30, !dbg !108

; <label>:29                                      ; preds = %26, %19
  br label %36, !dbg !110

; <label>:30                                      ; preds = %26
  %31 = load i32* %recvResult, align 4, !dbg !112
  %32 = sext i32 %31 to i64, !dbg !112
  %33 = getelementptr inbounds [14 x i8]* %inputBuffer, i32 0, i64 %32, !dbg !112
  store i8 0, i8* %33, align 1, !dbg !112
  %34 = getelementptr inbounds [14 x i8]* %inputBuffer, i32 0, i32 0, !dbg !113
  %35 = call i32 @atoi(i8* %34) #7, !dbg !113
  store i32 %35, i32* %data, align 4, !dbg !113
  br label %36, !dbg !114

; <label>:36                                      ; preds = %30, %29, %18, %5
  %37 = load i32* %connectSocket, align 4, !dbg !115
  %38 = icmp ne i32 %37, -1, !dbg !115
  br i1 %38, label %39, label %42, !dbg !115

; <label>:39                                      ; preds = %36
  %40 = load i32* %connectSocket, align 4, !dbg !117
  %41 = call i32 @close(i32 %40), !dbg !117
  br label %42, !dbg !119

; <label>:42                                      ; preds = %39, %36
  call void @llvm.dbg.declare(metadata !{i32* %i}, metadata !120), !dbg !122
  call void @llvm.dbg.declare(metadata !{[10 x i32]* %buffer}, metadata !123), !dbg !127
  %43 = bitcast [10 x i32]* %buffer to i8*, !dbg !127
  call void @llvm.memset.p0i8.i64(i8* %43, i8 0, i64 40, i32 16, i1 false), !dbg !127
  %44 = load i32* %data, align 4, !dbg !128
  %45 = icmp sge i32 %44, 0, !dbg !128
  br i1 %45, label %46, label %62, !dbg !128

; <label>:46                                      ; preds = %42
  %47 = load i32* %data, align 4, !dbg !130
  %48 = sext i32 %47 to i64, !dbg !130
  %49 = getelementptr inbounds [10 x i32]* %buffer, i32 0, i64 %48, !dbg !130
  store i32 1, i32* %49, align 4, !dbg !130
  store i32 0, i32* %i, align 4, !dbg !132
  br label %50, !dbg !132

; <label>:50                                      ; preds = %58, %46
  %51 = load i32* %i, align 4, !dbg !132
  %52 = icmp slt i32 %51, 10, !dbg !132
  br i1 %52, label %53, label %61, !dbg !132

; <label>:53                                      ; preds = %50
  %54 = load i32* %i, align 4, !dbg !134
  %55 = sext i32 %54 to i64, !dbg !134
  %56 = getelementptr inbounds [10 x i32]* %buffer, i32 0, i64 %55, !dbg !134
  %57 = load i32* %56, align 4, !dbg !134
  call void @printIntLine(i32 %57), !dbg !134
  br label %58, !dbg !136

; <label>:58                                      ; preds = %53
  %59 = load i32* %i, align 4, !dbg !132
  %60 = add nsw i32 %59, 1, !dbg !132
  store i32 %60, i32* %i, align 4, !dbg !132
  br label %50, !dbg !132

; <label>:61                                      ; preds = %50
  br label %63, !dbg !137

; <label>:62                                      ; preds = %42
  call void @printLine(i8* getelementptr inbounds ([32 x i8]* @.str1, i32 0, i32 0)), !dbg !138
  br label %63

; <label>:63                                      ; preds = %62, %61
  ret void, !dbg !140
}

; Function Attrs: nounwind readnone
declare void @llvm.dbg.declare(metadata, metadata) #1

; Function Attrs: nounwind
declare i32 @socket(i32, i32, i32) #2

; Function Attrs: nounwind
declare void @llvm.memset.p0i8.i64(i8* nocapture, i8, i64, i32, i1) #3

; Function Attrs: nounwind
declare i32 @inet_addr(i8*) #2

; Function Attrs: nounwind readnone
declare zeroext i16 @htons(i16 zeroext) #4

declare i32 @connect(i32, %struct.sockaddr*, i32) #5

declare i64 @recv(i32, i8*, i64, i32) #5

; Function Attrs: nounwind readonly
declare i32 @atoi(i8*) #6

declare i32 @close(i32) #5

declare void @printIntLine(i32) #5

declare void @printLine(i8*) #5

; Function Attrs: nounwind uwtable
define void @CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_good() #0 {
  call void @goodG2B(), !dbg !141
  call void @goodB2G(), !dbg !142
  ret void, !dbg !143
}

; Function Attrs: nounwind uwtable
define internal void @goodB2G() #0 {
  %data = alloca i32, align 4
  %recvResult = alloca i32, align 4
  %service = alloca %struct.sockaddr_in, align 4
  %connectSocket = alloca i32, align 4
  %inputBuffer = alloca [14 x i8], align 1
  %i = alloca i32, align 4
  %buffer = alloca [10 x i32], align 16
  call void @llvm.dbg.declare(metadata !{i32* %data}, metadata !144), !dbg !145
  store i32 -1, i32* %data, align 4, !dbg !146
  call void @llvm.dbg.declare(metadata !{i32* %recvResult}, metadata !147), !dbg !149
  call void @llvm.dbg.declare(metadata !{%struct.sockaddr_in* %service}, metadata !150), !dbg !151
  call void @llvm.dbg.declare(metadata !{i32* %connectSocket}, metadata !152), !dbg !153
  store i32 -1, i32* %connectSocket, align 4, !dbg !153
  call void @llvm.dbg.declare(metadata !{[14 x i8]* %inputBuffer}, metadata !154), !dbg !155
  br label %1, !dbg !156

; <label>:1                                       ; preds = %0
  %2 = call i32 @socket(i32 2, i32 1, i32 6) #3, !dbg !157
  store i32 %2, i32* %connectSocket, align 4, !dbg !157
  %3 = load i32* %connectSocket, align 4, !dbg !159
  %4 = icmp eq i32 %3, -1, !dbg !159
  br i1 %4, label %5, label %6, !dbg !159

; <label>:5                                       ; preds = %1
  br label %36, !dbg !161

; <label>:6                                       ; preds = %1
  %7 = bitcast %struct.sockaddr_in* %service to i8*, !dbg !163
  call void @llvm.memset.p0i8.i64(i8* %7, i8 0, i64 16, i32 4, i1 false), !dbg !163
  %8 = getelementptr inbounds %struct.sockaddr_in* %service, i32 0, i32 0, !dbg !164
  store i16 2, i16* %8, align 2, !dbg !164
  %9 = call i32 @inet_addr(i8* getelementptr inbounds ([10 x i8]* @.str, i32 0, i32 0)) #3, !dbg !165
  %10 = getelementptr inbounds %struct.sockaddr_in* %service, i32 0, i32 2, !dbg !165
  %11 = getelementptr inbounds %struct.in_addr* %10, i32 0, i32 0, !dbg !165
  store i32 %9, i32* %11, align 4, !dbg !165
  %12 = call zeroext i16 @htons(i16 zeroext 27015) #1, !dbg !166
  %13 = getelementptr inbounds %struct.sockaddr_in* %service, i32 0, i32 1, !dbg !166
  store i16 %12, i16* %13, align 2, !dbg !166
  %14 = load i32* %connectSocket, align 4, !dbg !167
  %15 = bitcast %struct.sockaddr_in* %service to %struct.sockaddr*, !dbg !167
  %16 = call i32 @connect(i32 %14, %struct.sockaddr* %15, i32 16), !dbg !167
  %17 = icmp eq i32 %16, -1, !dbg !167
  br i1 %17, label %18, label %19, !dbg !167

; <label>:18                                      ; preds = %6
  br label %36, !dbg !169

; <label>:19                                      ; preds = %6
  %20 = load i32* %connectSocket, align 4, !dbg !171
  %21 = getelementptr inbounds [14 x i8]* %inputBuffer, i32 0, i32 0, !dbg !171
  %22 = call i64 @recv(i32 %20, i8* %21, i64 13, i32 0), !dbg !171
  %23 = trunc i64 %22 to i32, !dbg !171
  store i32 %23, i32* %recvResult, align 4, !dbg !171
  %24 = load i32* %recvResult, align 4, !dbg !172
  %25 = icmp eq i32 %24, -1, !dbg !172
  br i1 %25, label %29, label %26, !dbg !172

; <label>:26                                      ; preds = %19
  %27 = load i32* %recvResult, align 4, !dbg !172
  %28 = icmp eq i32 %27, 0, !dbg !172
  br i1 %28, label %29, label %30, !dbg !172

; <label>:29                                      ; preds = %26, %19
  br label %36, !dbg !174

; <label>:30                                      ; preds = %26
  %31 = load i32* %recvResult, align 4, !dbg !176
  %32 = sext i32 %31 to i64, !dbg !176
  %33 = getelementptr inbounds [14 x i8]* %inputBuffer, i32 0, i64 %32, !dbg !176
  store i8 0, i8* %33, align 1, !dbg !176
  %34 = getelementptr inbounds [14 x i8]* %inputBuffer, i32 0, i32 0, !dbg !177
  %35 = call i32 @atoi(i8* %34) #7, !dbg !177
  store i32 %35, i32* %data, align 4, !dbg !177
  br label %36, !dbg !178

; <label>:36                                      ; preds = %30, %29, %18, %5
  %37 = load i32* %connectSocket, align 4, !dbg !179
  %38 = icmp ne i32 %37, -1, !dbg !179
  br i1 %38, label %39, label %42, !dbg !179

; <label>:39                                      ; preds = %36
  %40 = load i32* %connectSocket, align 4, !dbg !181
  %41 = call i32 @close(i32 %40), !dbg !181
  br label %42, !dbg !183

; <label>:42                                      ; preds = %39, %36
  call void @llvm.dbg.declare(metadata !{i32* %i}, metadata !184), !dbg !186
  call void @llvm.dbg.declare(metadata !{[10 x i32]* %buffer}, metadata !187), !dbg !188
  %43 = bitcast [10 x i32]* %buffer to i8*, !dbg !188
  call void @llvm.memset.p0i8.i64(i8* %43, i8 0, i64 40, i32 16, i1 false), !dbg !188
  %44 = load i32* %data, align 4, !dbg !189
  %45 = icmp sge i32 %44, 0, !dbg !189
  br i1 %45, label %46, label %65, !dbg !189

; <label>:46                                      ; preds = %42
  %47 = load i32* %data, align 4, !dbg !189
  %48 = icmp slt i32 %47, 10, !dbg !189
  br i1 %48, label %49, label %65, !dbg !189

; <label>:49                                      ; preds = %46
  %50 = load i32* %data, align 4, !dbg !191
  %51 = sext i32 %50 to i64, !dbg !191
  %52 = getelementptr inbounds [10 x i32]* %buffer, i32 0, i64 %51, !dbg !191
  store i32 1, i32* %52, align 4, !dbg !191
  store i32 0, i32* %i, align 4, !dbg !193
  br label %53, !dbg !193

; <label>:53                                      ; preds = %61, %49
  %54 = load i32* %i, align 4, !dbg !193
  %55 = icmp slt i32 %54, 10, !dbg !193
  br i1 %55, label %56, label %64, !dbg !193

; <label>:56                                      ; preds = %53
  %57 = load i32* %i, align 4, !dbg !195
  %58 = sext i32 %57 to i64, !dbg !195
  %59 = getelementptr inbounds [10 x i32]* %buffer, i32 0, i64 %58, !dbg !195
  %60 = load i32* %59, align 4, !dbg !195
  call void @printIntLine(i32 %60), !dbg !195
  br label %61, !dbg !197

; <label>:61                                      ; preds = %56
  %62 = load i32* %i, align 4, !dbg !193
  %63 = add nsw i32 %62, 1, !dbg !193
  store i32 %63, i32* %i, align 4, !dbg !193
  br label %53, !dbg !193

; <label>:64                                      ; preds = %53
  br label %66, !dbg !198

; <label>:65                                      ; preds = %46, %42
  call void @printLine(i8* getelementptr inbounds ([36 x i8]* @.str2, i32 0, i32 0)), !dbg !199
  br label %66

; <label>:66                                      ; preds = %65, %64
  ret void, !dbg !201
}

; Function Attrs: nounwind uwtable
define internal void @goodG2B() #0 {
  %data = alloca i32, align 4
  %i = alloca i32, align 4
  %buffer = alloca [10 x i32], align 16
  call void @llvm.dbg.declare(metadata !{i32* %data}, metadata !202), !dbg !203
  store i32 -1, i32* %data, align 4, !dbg !204
  store i32 7, i32* %data, align 4, !dbg !205
  call void @llvm.dbg.declare(metadata !{i32* %i}, metadata !206), !dbg !208
  call void @llvm.dbg.declare(metadata !{[10 x i32]* %buffer}, metadata !209), !dbg !210
  %1 = bitcast [10 x i32]* %buffer to i8*, !dbg !210
  call void @llvm.memset.p0i8.i64(i8* %1, i8 0, i64 40, i32 16, i1 false), !dbg !210
  %2 = load i32* %data, align 4, !dbg !211
  %3 = icmp sge i32 %2, 0, !dbg !211
  br i1 %3, label %4, label %20, !dbg !211

; <label>:4                                       ; preds = %0
  %5 = load i32* %data, align 4, !dbg !213
  %6 = sext i32 %5 to i64, !dbg !213
  %7 = getelementptr inbounds [10 x i32]* %buffer, i32 0, i64 %6, !dbg !213
  store i32 1, i32* %7, align 4, !dbg !213
  store i32 0, i32* %i, align 4, !dbg !215
  br label %8, !dbg !215

; <label>:8                                       ; preds = %16, %4
  %9 = load i32* %i, align 4, !dbg !215
  %10 = icmp slt i32 %9, 10, !dbg !215
  br i1 %10, label %11, label %19, !dbg !215

; <label>:11                                      ; preds = %8
  %12 = load i32* %i, align 4, !dbg !217
  %13 = sext i32 %12 to i64, !dbg !217
  %14 = getelementptr inbounds [10 x i32]* %buffer, i32 0, i64 %13, !dbg !217
  %15 = load i32* %14, align 4, !dbg !217
  call void @printIntLine(i32 %15), !dbg !217
  br label %16, !dbg !219

; <label>:16                                      ; preds = %11
  %17 = load i32* %i, align 4, !dbg !215
  %18 = add nsw i32 %17, 1, !dbg !215
  store i32 %18, i32* %i, align 4, !dbg !215
  br label %8, !dbg !215

; <label>:19                                      ; preds = %8
  br label %21, !dbg !220

; <label>:20                                      ; preds = %0
  call void @printLine(i8* getelementptr inbounds ([32 x i8]* @.str1, i32 0, i32 0)), !dbg !221
  br label %21

; <label>:21                                      ; preds = %20, %19
  ret void, !dbg !223
}

attributes #0 = { nounwind uwtable "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { nounwind readnone }
attributes #2 = { nounwind "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #3 = { nounwind }
attributes #4 = { nounwind readnone "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #5 = { "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #6 = { nounwind readonly "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #7 = { nounwind readonly }

!llvm.dbg.cu = !{!0}
!llvm.module.flags = !{!52, !53}
!llvm.ident = !{!54}

!0 = metadata !{i32 786449, metadata !1, i32 12, metadata !"clang version 3.4.2 (tags/RELEASE_34/dot2-final)", i1 false, metadata !"", i32 0, metadata !2, metadata !43, metadata !44, metadata !43, metadata !43, metadata !""} ; [ DW_TAG_compile_unit ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c] [DW_LANG_C99]
!1 = metadata !{metadata !"CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c", metadata !"/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01"}
!2 = metadata !{metadata !3, metadata !15}
!3 = metadata !{i32 786436, metadata !4, null, metadata !"__socket_type", i32 24, i64 32, i64 32, i32 0, i32 0, null, metadata !5, i32 0, null, null, null} ; [ DW_TAG_enumeration_type ] [__socket_type] [line 24, size 32, align 32, offset 0] [def] [from ]
!4 = metadata !{metadata !"/usr/include/bits/socket_type.h", metadata !"/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01"}
!5 = metadata !{metadata !6, metadata !7, metadata !8, metadata !9, metadata !10, metadata !11, metadata !12, metadata !13, metadata !14}
!6 = metadata !{i32 786472, metadata !"SOCK_STREAM", i64 1} ; [ DW_TAG_enumerator ] [SOCK_STREAM :: 1]
!7 = metadata !{i32 786472, metadata !"SOCK_DGRAM", i64 2} ; [ DW_TAG_enumerator ] [SOCK_DGRAM :: 2]
!8 = metadata !{i32 786472, metadata !"SOCK_RAW", i64 3} ; [ DW_TAG_enumerator ] [SOCK_RAW :: 3]
!9 = metadata !{i32 786472, metadata !"SOCK_RDM", i64 4} ; [ DW_TAG_enumerator ] [SOCK_RDM :: 4]
!10 = metadata !{i32 786472, metadata !"SOCK_SEQPACKET", i64 5} ; [ DW_TAG_enumerator ] [SOCK_SEQPACKET :: 5]
!11 = metadata !{i32 786472, metadata !"SOCK_DCCP", i64 6} ; [ DW_TAG_enumerator ] [SOCK_DCCP :: 6]
!12 = metadata !{i32 786472, metadata !"SOCK_PACKET", i64 10} ; [ DW_TAG_enumerator ] [SOCK_PACKET :: 10]
!13 = metadata !{i32 786472, metadata !"SOCK_CLOEXEC", i64 524288} ; [ DW_TAG_enumerator ] [SOCK_CLOEXEC :: 524288]
!14 = metadata !{i32 786472, metadata !"SOCK_NONBLOCK", i64 2048} ; [ DW_TAG_enumerator ] [SOCK_NONBLOCK :: 2048]
!15 = metadata !{i32 786436, metadata !16, null, metadata !"", i32 41, i64 32, i64 32, i32 0, i32 0, null, metadata !17, i32 0, null, null, null} ; [ DW_TAG_enumeration_type ] [line 41, size 32, align 32, offset 0] [def] [from ]
!16 = metadata !{metadata !"/usr/include/netinet/in.h", metadata !"/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01"}
!17 = metadata !{metadata !18, metadata !19, metadata !20, metadata !21, metadata !22, metadata !23, metadata !24, metadata !25, metadata !26, metadata !27, metadata !28, metadata !29, metadata !30, metadata !31, metadata !32, metadata !33, metadata !34, metadata !35, metadata !36, metadata !37, metadata !38, metadata !39, metadata !40, metadata !41, metadata !42}
!18 = metadata !{i32 786472, metadata !"IPPROTO_IP", i64 0} ; [ DW_TAG_enumerator ] [IPPROTO_IP :: 0]
!19 = metadata !{i32 786472, metadata !"IPPROTO_ICMP", i64 1} ; [ DW_TAG_enumerator ] [IPPROTO_ICMP :: 1]
!20 = metadata !{i32 786472, metadata !"IPPROTO_IGMP", i64 2} ; [ DW_TAG_enumerator ] [IPPROTO_IGMP :: 2]
!21 = metadata !{i32 786472, metadata !"IPPROTO_IPIP", i64 4} ; [ DW_TAG_enumerator ] [IPPROTO_IPIP :: 4]
!22 = metadata !{i32 786472, metadata !"IPPROTO_TCP", i64 6} ; [ DW_TAG_enumerator ] [IPPROTO_TCP :: 6]
!23 = metadata !{i32 786472, metadata !"IPPROTO_EGP", i64 8} ; [ DW_TAG_enumerator ] [IPPROTO_EGP :: 8]
!24 = metadata !{i32 786472, metadata !"IPPROTO_PUP", i64 12} ; [ DW_TAG_enumerator ] [IPPROTO_PUP :: 12]
!25 = metadata !{i32 786472, metadata !"IPPROTO_UDP", i64 17} ; [ DW_TAG_enumerator ] [IPPROTO_UDP :: 17]
!26 = metadata !{i32 786472, metadata !"IPPROTO_IDP", i64 22} ; [ DW_TAG_enumerator ] [IPPROTO_IDP :: 22]
!27 = metadata !{i32 786472, metadata !"IPPROTO_TP", i64 29} ; [ DW_TAG_enumerator ] [IPPROTO_TP :: 29]
!28 = metadata !{i32 786472, metadata !"IPPROTO_DCCP", i64 33} ; [ DW_TAG_enumerator ] [IPPROTO_DCCP :: 33]
!29 = metadata !{i32 786472, metadata !"IPPROTO_IPV6", i64 41} ; [ DW_TAG_enumerator ] [IPPROTO_IPV6 :: 41]
!30 = metadata !{i32 786472, metadata !"IPPROTO_RSVP", i64 46} ; [ DW_TAG_enumerator ] [IPPROTO_RSVP :: 46]
!31 = metadata !{i32 786472, metadata !"IPPROTO_GRE", i64 47} ; [ DW_TAG_enumerator ] [IPPROTO_GRE :: 47]
!32 = metadata !{i32 786472, metadata !"IPPROTO_ESP", i64 50} ; [ DW_TAG_enumerator ] [IPPROTO_ESP :: 50]
!33 = metadata !{i32 786472, metadata !"IPPROTO_AH", i64 51} ; [ DW_TAG_enumerator ] [IPPROTO_AH :: 51]
!34 = metadata !{i32 786472, metadata !"IPPROTO_MTP", i64 92} ; [ DW_TAG_enumerator ] [IPPROTO_MTP :: 92]
!35 = metadata !{i32 786472, metadata !"IPPROTO_BEETPH", i64 94} ; [ DW_TAG_enumerator ] [IPPROTO_BEETPH :: 94]
!36 = metadata !{i32 786472, metadata !"IPPROTO_ENCAP", i64 98} ; [ DW_TAG_enumerator ] [IPPROTO_ENCAP :: 98]
!37 = metadata !{i32 786472, metadata !"IPPROTO_PIM", i64 103} ; [ DW_TAG_enumerator ] [IPPROTO_PIM :: 103]
!38 = metadata !{i32 786472, metadata !"IPPROTO_COMP", i64 108} ; [ DW_TAG_enumerator ] [IPPROTO_COMP :: 108]
!39 = metadata !{i32 786472, metadata !"IPPROTO_SCTP", i64 132} ; [ DW_TAG_enumerator ] [IPPROTO_SCTP :: 132]
!40 = metadata !{i32 786472, metadata !"IPPROTO_UDPLITE", i64 136} ; [ DW_TAG_enumerator ] [IPPROTO_UDPLITE :: 136]
!41 = metadata !{i32 786472, metadata !"IPPROTO_RAW", i64 255} ; [ DW_TAG_enumerator ] [IPPROTO_RAW :: 255]
!42 = metadata !{i32 786472, metadata !"IPPROTO_MAX", i64 256} ; [ DW_TAG_enumerator ] [IPPROTO_MAX :: 256]
!43 = metadata !{i32 0}
!44 = metadata !{metadata !45, metadata !49, metadata !50, metadata !51}
!45 = metadata !{i32 786478, metadata !1, metadata !46, metadata !"CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_bad", metadata !"CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_bad", metadata !"", i32 44, metadata !47, i1 false, i1 true, i32 0, i32 0, null, i32 0, i1 false, void ()* @CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_bad, null, null, metadata !43, i32 45} ; [ DW_TAG_subprogram ] [line 44] [def] [scope 45] [CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_bad]
!46 = metadata !{i32 786473, metadata !1}         ; [ DW_TAG_file_type ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!47 = metadata !{i32 786453, i32 0, null, metadata !"", i32 0, i64 0, i64 0, i64 0, i32 0, null, metadata !48, i32 0, null, null, null} ; [ DW_TAG_subroutine_type ] [line 0, size 0, align 0, offset 0] [from ]
!48 = metadata !{null}
!49 = metadata !{i32 786478, metadata !1, metadata !46, metadata !"CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_good", metadata !"CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_good", metadata !"", i32 242, metadata !47, i1 false, i1 true, i32 0, i32 0, null, i32 0, i1 false, void ()* @CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_good, null, null, metadata !43, i32 243} ; [ DW_TAG_subprogram ] [line 242] [def] [scope 243] [CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01_good]
!50 = metadata !{i32 786478, metadata !1, metadata !46, metadata !"goodB2G", metadata !"goodB2G", metadata !"", i32 161, metadata !47, i1 true, i1 true, i32 0, i32 0, null, i32 0, i1 false, void ()* @goodB2G, null, null, metadata !43, i32 162} ; [ DW_TAG_subprogram ] [line 161] [local] [def] [scope 162] [goodB2G]
!51 = metadata !{i32 786478, metadata !1, metadata !46, metadata !"goodG2B", metadata !"goodG2B", metadata !"", i32 131, metadata !47, i1 true, i1 true, i32 0, i32 0, null, i32 0, i1 false, void ()* @goodG2B, null, null, metadata !43, i32 132} ; [ DW_TAG_subprogram ] [line 131] [local] [def] [scope 132] [goodG2B]
!52 = metadata !{i32 2, metadata !"Dwarf Version", i32 4}
!53 = metadata !{i32 1, metadata !"Debug Info Version", i32 1}
!54 = metadata !{metadata !"clang version 3.4.2 (tags/RELEASE_34/dot2-final)"}
!55 = metadata !{i32 786688, metadata !45, metadata !"data", metadata !46, i32 46, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [data] [line 46]
!56 = metadata !{i32 786468, null, null, metadata !"int", i32 0, i64 32, i64 32, i64 0, i32 0, i32 5} ; [ DW_TAG_base_type ] [int] [line 0, size 32, align 32, offset 0, enc DW_ATE_signed]
!57 = metadata !{i32 46, i32 0, metadata !45, null}
!58 = metadata !{i32 48, i32 0, metadata !45, null}
!59 = metadata !{i32 786688, metadata !60, metadata !"recvResult", metadata !46, i32 54, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [recvResult] [line 54]
!60 = metadata !{i32 786443, metadata !1, metadata !45, i32 49, i32 0, i32 0} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!61 = metadata !{i32 54, i32 0, metadata !60, null}
!62 = metadata !{i32 786688, metadata !60, metadata !"service", metadata !46, i32 55, metadata !63, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [service] [line 55]
!63 = metadata !{i32 786451, metadata !16, null, metadata !"sockaddr_in", i32 238, i64 128, i64 32, i32 0, i32 0, null, metadata !64, i32 0, null, null, null} ; [ DW_TAG_structure_type ] [sockaddr_in] [line 238, size 128, align 32, offset 0] [def] [from ]
!64 = metadata !{metadata !65, metadata !68, metadata !71, metadata !78}
!65 = metadata !{i32 786445, metadata !16, metadata !63, metadata !"sin_family", i32 240, i64 16, i64 16, i64 0, i32 0, metadata !66} ; [ DW_TAG_member ] [sin_family] [line 240, size 16, align 16, offset 0] [from sa_family_t]
!66 = metadata !{i32 786454, metadata !16, null, metadata !"sa_family_t", i32 28, i64 0, i64 0, i64 0, i32 0, metadata !67} ; [ DW_TAG_typedef ] [sa_family_t] [line 28, size 0, align 0, offset 0] [from unsigned short]
!67 = metadata !{i32 786468, null, null, metadata !"unsigned short", i32 0, i64 16, i64 16, i64 0, i32 0, i32 7} ; [ DW_TAG_base_type ] [unsigned short] [line 0, size 16, align 16, offset 0, enc DW_ATE_unsigned]
!68 = metadata !{i32 786445, metadata !16, metadata !63, metadata !"sin_port", i32 241, i64 16, i64 16, i64 16, i32 0, metadata !69} ; [ DW_TAG_member ] [sin_port] [line 241, size 16, align 16, offset 16] [from in_port_t]
!69 = metadata !{i32 786454, metadata !16, null, metadata !"in_port_t", i32 118, i64 0, i64 0, i64 0, i32 0, metadata !70} ; [ DW_TAG_typedef ] [in_port_t] [line 118, size 0, align 0, offset 0] [from uint16_t]
!70 = metadata !{i32 786454, metadata !16, null, metadata !"uint16_t", i32 49, i64 0, i64 0, i64 0, i32 0, metadata !67} ; [ DW_TAG_typedef ] [uint16_t] [line 49, size 0, align 0, offset 0] [from unsigned short]
!71 = metadata !{i32 786445, metadata !16, metadata !63, metadata !"sin_addr", i32 242, i64 32, i64 32, i64 32, i32 0, metadata !72} ; [ DW_TAG_member ] [sin_addr] [line 242, size 32, align 32, offset 32] [from in_addr]
!72 = metadata !{i32 786451, metadata !16, null, metadata !"in_addr", i32 32, i64 32, i64 32, i32 0, i32 0, null, metadata !73, i32 0, null, null, null} ; [ DW_TAG_structure_type ] [in_addr] [line 32, size 32, align 32, offset 0] [def] [from ]
!73 = metadata !{metadata !74}
!74 = metadata !{i32 786445, metadata !16, metadata !72, metadata !"s_addr", i32 34, i64 32, i64 32, i64 0, i32 0, metadata !75} ; [ DW_TAG_member ] [s_addr] [line 34, size 32, align 32, offset 0] [from in_addr_t]
!75 = metadata !{i32 786454, metadata !16, null, metadata !"in_addr_t", i32 31, i64 0, i64 0, i64 0, i32 0, metadata !76} ; [ DW_TAG_typedef ] [in_addr_t] [line 31, size 0, align 0, offset 0] [from uint32_t]
!76 = metadata !{i32 786454, metadata !16, null, metadata !"uint32_t", i32 51, i64 0, i64 0, i64 0, i32 0, metadata !77} ; [ DW_TAG_typedef ] [uint32_t] [line 51, size 0, align 0, offset 0] [from unsigned int]
!77 = metadata !{i32 786468, null, null, metadata !"unsigned int", i32 0, i64 32, i64 32, i64 0, i32 0, i32 7} ; [ DW_TAG_base_type ] [unsigned int] [line 0, size 32, align 32, offset 0, enc DW_ATE_unsigned]
!78 = metadata !{i32 786445, metadata !16, metadata !63, metadata !"sin_zero", i32 245, i64 64, i64 8, i64 64, i32 0, metadata !79} ; [ DW_TAG_member ] [sin_zero] [line 245, size 64, align 8, offset 64] [from ]
!79 = metadata !{i32 786433, null, null, metadata !"", i32 0, i64 64, i64 8, i32 0, i32 0, metadata !80, metadata !81, i32 0, null, null, null} ; [ DW_TAG_array_type ] [line 0, size 64, align 8, offset 0] [from unsigned char]
!80 = metadata !{i32 786468, null, null, metadata !"unsigned char", i32 0, i64 8, i64 8, i64 0, i32 0, i32 8} ; [ DW_TAG_base_type ] [unsigned char] [line 0, size 8, align 8, offset 0, enc DW_ATE_unsigned_char]
!81 = metadata !{metadata !82}
!82 = metadata !{i32 786465, i64 0, i64 8}        ; [ DW_TAG_subrange_type ] [0, 7]
!83 = metadata !{i32 55, i32 0, metadata !60, null}
!84 = metadata !{i32 786688, metadata !60, metadata !"connectSocket", metadata !46, i32 56, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [connectSocket] [line 56]
!85 = metadata !{i32 56, i32 0, metadata !60, null}
!86 = metadata !{i32 786688, metadata !60, metadata !"inputBuffer", metadata !46, i32 57, metadata !87, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [inputBuffer] [line 57]
!87 = metadata !{i32 786433, null, null, metadata !"", i32 0, i64 112, i64 8, i32 0, i32 0, metadata !88, metadata !89, i32 0, null, null, null} ; [ DW_TAG_array_type ] [line 0, size 112, align 8, offset 0] [from char]
!88 = metadata !{i32 786468, null, null, metadata !"char", i32 0, i64 8, i64 8, i64 0, i32 0, i32 6} ; [ DW_TAG_base_type ] [char] [line 0, size 8, align 8, offset 0, enc DW_ATE_signed_char]
!89 = metadata !{metadata !90}
!90 = metadata !{i32 786465, i64 0, i64 14}       ; [ DW_TAG_subrange_type ] [0, 13]
!91 = metadata !{i32 57, i32 0, metadata !60, null}
!92 = metadata !{i32 58, i32 0, metadata !60, null} ; [ DW_TAG_imported_module ]
!93 = metadata !{i32 68, i32 0, metadata !94, null}
!94 = metadata !{i32 786443, metadata !1, metadata !60, i32 59, i32 0, i32 1} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!95 = metadata !{i32 69, i32 0, metadata !96, null}
!96 = metadata !{i32 786443, metadata !1, metadata !94, i32 69, i32 0, i32 2} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!97 = metadata !{i32 71, i32 0, metadata !98, null}
!98 = metadata !{i32 786443, metadata !1, metadata !96, i32 70, i32 0, i32 3} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!99 = metadata !{i32 73, i32 0, metadata !94, null}
!100 = metadata !{i32 74, i32 0, metadata !94, null}
!101 = metadata !{i32 75, i32 0, metadata !94, null}
!102 = metadata !{i32 76, i32 0, metadata !94, null}
!103 = metadata !{i32 77, i32 0, metadata !104, null}
!104 = metadata !{i32 786443, metadata !1, metadata !94, i32 77, i32 0, i32 4} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!105 = metadata !{i32 79, i32 0, metadata !106, null}
!106 = metadata !{i32 786443, metadata !1, metadata !104, i32 78, i32 0, i32 5} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!107 = metadata !{i32 83, i32 0, metadata !94, null}
!108 = metadata !{i32 84, i32 0, metadata !109, null}
!109 = metadata !{i32 786443, metadata !1, metadata !94, i32 84, i32 0, i32 6} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!110 = metadata !{i32 86, i32 0, metadata !111, null}
!111 = metadata !{i32 786443, metadata !1, metadata !109, i32 85, i32 0, i32 7} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!112 = metadata !{i32 89, i32 0, metadata !94, null}
!113 = metadata !{i32 91, i32 0, metadata !94, null}
!114 = metadata !{i32 92, i32 0, metadata !94, null}
!115 = metadata !{i32 94, i32 0, metadata !116, null}
!116 = metadata !{i32 786443, metadata !1, metadata !60, i32 94, i32 0, i32 8} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!117 = metadata !{i32 96, i32 0, metadata !118, null}
!118 = metadata !{i32 786443, metadata !1, metadata !116, i32 95, i32 0, i32 9} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!119 = metadata !{i32 97, i32 0, metadata !118, null}
!120 = metadata !{i32 786688, metadata !121, metadata !"i", metadata !46, i32 106, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [i] [line 106]
!121 = metadata !{i32 786443, metadata !1, metadata !45, i32 105, i32 0, i32 10} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!122 = metadata !{i32 106, i32 0, metadata !121, null}
!123 = metadata !{i32 786688, metadata !121, metadata !"buffer", metadata !46, i32 107, metadata !124, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [buffer] [line 107]
!124 = metadata !{i32 786433, null, null, metadata !"", i32 0, i64 320, i64 32, i32 0, i32 0, metadata !56, metadata !125, i32 0, null, null, null} ; [ DW_TAG_array_type ] [line 0, size 320, align 32, offset 0] [from int]
!125 = metadata !{metadata !126}
!126 = metadata !{i32 786465, i64 0, i64 10}      ; [ DW_TAG_subrange_type ] [0, 9]
!127 = metadata !{i32 107, i32 0, metadata !121, null}
!128 = metadata !{i32 110, i32 0, metadata !129, null}
!129 = metadata !{i32 786443, metadata !1, metadata !121, i32 110, i32 0, i32 11} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!130 = metadata !{i32 112, i32 0, metadata !131, null}
!131 = metadata !{i32 786443, metadata !1, metadata !129, i32 111, i32 0, i32 12} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!132 = metadata !{i32 114, i32 0, metadata !133, null}
!133 = metadata !{i32 786443, metadata !1, metadata !131, i32 114, i32 0, i32 13} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!134 = metadata !{i32 116, i32 0, metadata !135, null}
!135 = metadata !{i32 786443, metadata !1, metadata !133, i32 115, i32 0, i32 14} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!136 = metadata !{i32 117, i32 0, metadata !135, null}
!137 = metadata !{i32 118, i32 0, metadata !131, null}
!138 = metadata !{i32 121, i32 0, metadata !139, null}
!139 = metadata !{i32 786443, metadata !1, metadata !129, i32 120, i32 0, i32 15} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!140 = metadata !{i32 124, i32 0, metadata !45, null}
!141 = metadata !{i32 244, i32 0, metadata !49, null}
!142 = metadata !{i32 245, i32 0, metadata !49, null}
!143 = metadata !{i32 246, i32 0, metadata !49, null}
!144 = metadata !{i32 786688, metadata !50, metadata !"data", metadata !46, i32 163, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [data] [line 163]
!145 = metadata !{i32 163, i32 0, metadata !50, null}
!146 = metadata !{i32 165, i32 0, metadata !50, null}
!147 = metadata !{i32 786688, metadata !148, metadata !"recvResult", metadata !46, i32 171, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [recvResult] [line 171]
!148 = metadata !{i32 786443, metadata !1, metadata !50, i32 166, i32 0, i32 16} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!149 = metadata !{i32 171, i32 0, metadata !148, null}
!150 = metadata !{i32 786688, metadata !148, metadata !"service", metadata !46, i32 172, metadata !63, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [service] [line 172]
!151 = metadata !{i32 172, i32 0, metadata !148, null}
!152 = metadata !{i32 786688, metadata !148, metadata !"connectSocket", metadata !46, i32 173, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [connectSocket] [line 173]
!153 = metadata !{i32 173, i32 0, metadata !148, null}
!154 = metadata !{i32 786688, metadata !148, metadata !"inputBuffer", metadata !46, i32 174, metadata !87, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [inputBuffer] [line 174]
!155 = metadata !{i32 174, i32 0, metadata !148, null}
!156 = metadata !{i32 175, i32 0, metadata !148, null}
!157 = metadata !{i32 185, i32 0, metadata !158, null}
!158 = metadata !{i32 786443, metadata !1, metadata !148, i32 176, i32 0, i32 17} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!159 = metadata !{i32 186, i32 0, metadata !160, null}
!160 = metadata !{i32 786443, metadata !1, metadata !158, i32 186, i32 0, i32 18} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!161 = metadata !{i32 188, i32 0, metadata !162, null}
!162 = metadata !{i32 786443, metadata !1, metadata !160, i32 187, i32 0, i32 19} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!163 = metadata !{i32 190, i32 0, metadata !158, null}
!164 = metadata !{i32 191, i32 0, metadata !158, null}
!165 = metadata !{i32 192, i32 0, metadata !158, null}
!166 = metadata !{i32 193, i32 0, metadata !158, null}
!167 = metadata !{i32 194, i32 0, metadata !168, null}
!168 = metadata !{i32 786443, metadata !1, metadata !158, i32 194, i32 0, i32 20} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!169 = metadata !{i32 196, i32 0, metadata !170, null}
!170 = metadata !{i32 786443, metadata !1, metadata !168, i32 195, i32 0, i32 21} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!171 = metadata !{i32 200, i32 0, metadata !158, null}
!172 = metadata !{i32 201, i32 0, metadata !173, null}
!173 = metadata !{i32 786443, metadata !1, metadata !158, i32 201, i32 0, i32 22} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!174 = metadata !{i32 203, i32 0, metadata !175, null}
!175 = metadata !{i32 786443, metadata !1, metadata !173, i32 202, i32 0, i32 23} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!176 = metadata !{i32 206, i32 0, metadata !158, null}
!177 = metadata !{i32 208, i32 0, metadata !158, null}
!178 = metadata !{i32 209, i32 0, metadata !158, null}
!179 = metadata !{i32 211, i32 0, metadata !180, null}
!180 = metadata !{i32 786443, metadata !1, metadata !148, i32 211, i32 0, i32 24} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!181 = metadata !{i32 213, i32 0, metadata !182, null}
!182 = metadata !{i32 786443, metadata !1, metadata !180, i32 212, i32 0, i32 25} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!183 = metadata !{i32 214, i32 0, metadata !182, null}
!184 = metadata !{i32 786688, metadata !185, metadata !"i", metadata !46, i32 223, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [i] [line 223]
!185 = metadata !{i32 786443, metadata !1, metadata !50, i32 222, i32 0, i32 26} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!186 = metadata !{i32 223, i32 0, metadata !185, null}
!187 = metadata !{i32 786688, metadata !185, metadata !"buffer", metadata !46, i32 224, metadata !124, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [buffer] [line 224]
!188 = metadata !{i32 224, i32 0, metadata !185, null}
!189 = metadata !{i32 226, i32 0, metadata !190, null}
!190 = metadata !{i32 786443, metadata !1, metadata !185, i32 226, i32 0, i32 27} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!191 = metadata !{i32 228, i32 0, metadata !192, null}
!192 = metadata !{i32 786443, metadata !1, metadata !190, i32 227, i32 0, i32 28} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!193 = metadata !{i32 230, i32 0, metadata !194, null}
!194 = metadata !{i32 786443, metadata !1, metadata !192, i32 230, i32 0, i32 29} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!195 = metadata !{i32 232, i32 0, metadata !196, null}
!196 = metadata !{i32 786443, metadata !1, metadata !194, i32 231, i32 0, i32 30} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!197 = metadata !{i32 233, i32 0, metadata !196, null}
!198 = metadata !{i32 234, i32 0, metadata !192, null}
!199 = metadata !{i32 237, i32 0, metadata !200, null}
!200 = metadata !{i32 786443, metadata !1, metadata !190, i32 236, i32 0, i32 31} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!201 = metadata !{i32 240, i32 0, metadata !50, null}
!202 = metadata !{i32 786688, metadata !51, metadata !"data", metadata !46, i32 133, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [data] [line 133]
!203 = metadata !{i32 133, i32 0, metadata !51, null}
!204 = metadata !{i32 135, i32 0, metadata !51, null}
!205 = metadata !{i32 138, i32 0, metadata !51, null}
!206 = metadata !{i32 786688, metadata !207, metadata !"i", metadata !46, i32 140, metadata !56, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [i] [line 140]
!207 = metadata !{i32 786443, metadata !1, metadata !51, i32 139, i32 0, i32 32} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!208 = metadata !{i32 140, i32 0, metadata !207, null}
!209 = metadata !{i32 786688, metadata !207, metadata !"buffer", metadata !46, i32 141, metadata !124, i32 0, i32 0} ; [ DW_TAG_auto_variable ] [buffer] [line 141]
!210 = metadata !{i32 141, i32 0, metadata !207, null}
!211 = metadata !{i32 144, i32 0, metadata !212, null}
!212 = metadata !{i32 786443, metadata !1, metadata !207, i32 144, i32 0, i32 33} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!213 = metadata !{i32 146, i32 0, metadata !214, null}
!214 = metadata !{i32 786443, metadata !1, metadata !212, i32 145, i32 0, i32 34} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!215 = metadata !{i32 148, i32 0, metadata !216, null}
!216 = metadata !{i32 786443, metadata !1, metadata !214, i32 148, i32 0, i32 35} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!217 = metadata !{i32 150, i32 0, metadata !218, null}
!218 = metadata !{i32 786443, metadata !1, metadata !216, i32 149, i32 0, i32 36} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!219 = metadata !{i32 151, i32 0, metadata !218, null}
!220 = metadata !{i32 152, i32 0, metadata !214, null}
!221 = metadata !{i32 155, i32 0, metadata !222, null}
!222 = metadata !{i32 786443, metadata !1, metadata !212, i32 154, i32 0, i32 37} ; [ DW_TAG_lexical_block ] [/home/sqy5331/project/AST-transformer/data/C/testcases/CWE121_Stack_Based_Buffer_Overflow/s01/CWE121_Stack_Based_Buffer_Overflow__CWE129_connect_socket_01.c]
!223 = metadata !{i32 158, i32 0, metadata !51, null}
