/*
 * Copyright (c) 2018, Penn State University
 *
 * This software was developed by Zhilong Wang <njuwangzhilong@gmail.com>
 * at Penn State University, State College, PA, USA, in September 2018.
 *
 *
 *==-------------- Support.cpp --------------==//
 * This file contains some functions for analyizing.  
 */
#include "DFAIdentifier/Support.h"
#include "DFAIdentifier/Struct.h"
#include "DFAIdentifier/Switch.h"
#include "LiveAnalysis/LiveAnalysis.h"

#include <fstream>
#define MAXDEEP 3 //max depth when tracking the load instructions of cmp
#define DEF_USE_DBG 
#define STRUCT_DBG 
//#define INEV_BB
/* Get instruction id.
 */
unsigned Support::getInstructionID(Instruction * I) 
{
	MDNode * Node = I->getMetadata("ins_id");
	if(!Node)
	{
		return -1;
	}
	assert(Node->getNumOperands() == 1);
	const Metadata * MD = Node->getOperand(0);
	if(auto *MDV = dyn_cast<ValueAsMetadata>(MD))
	{
		Value * V = MDV->getValue();
		ConstantInt * CI = dyn_cast<ConstantInt>(V);
		assert(CI);
		return CI->getZExtValue();
	}
	return -1;
}
/* Print function location in source code.
 */
void Support::printFunLocation(Function * Fun)
{
	for (Function::iterator BB = Fun->begin(); BB != Fun->end(); BB++) {
		for (BasicBlock::iterator II = BB->begin(); II != BB->end(); II++) {
			Instruction *I = &*II;
			const DILocation * DIL = I->getDebugLoc();
			if(!DIL){
				continue;
			}else{
				char pPath[200];
				string sFileName;
				if(DIL->getFilename().str().substr(0,5) == "/home"){
					sFileName = DIL->getFilename().str();
				}else{
					sFileName = DIL->getDirectory().str() + "/" + DIL->getFilename().str();
				}
				realpath(sFileName.c_str(), pPath);
				sFileName = string(sFileName);
				unsigned int numLine = DIL->getLine();
				errs() << "//---" << sFileName << ": " << numLine << "\n";
				return;
			}
		}
	}
	return;
}
/* Print the basic blick location in source code.
 */
void Support::printSrcCodeInfo(BasicBlock * BB)
{
	for (BasicBlock::iterator II = BB->begin(); II != BB->end(); II++) {
	    Instruction *I = &*II;
		const DILocation * DIL = I->getDebugLoc();
		if(!DIL){
			continue;
		}else{
			char pPath[200];
			string sFileName;
			if(DIL->getFilename().str().substr(0,5) == "/home"){
				sFileName = DIL->getFilename().str();
			}else{
				sFileName = DIL->getDirectory().str() + "/" + DIL->getFilename().str();
			}
			realpath(sFileName.c_str(), pPath);
			sFileName = string(sFileName);
			unsigned int numLine = DIL->getLine();
			errs() << "//---" << sFileName << ": " << numLine << "\n";
			break;
		}
	}
	return;
}
/* Print the instruction location in source code.
 */
void Support::printSrcCodeInfo(Instruction * Ins)
{

	Instruction *I = &*Ins;
	const DILocation * DIL = I->getDebugLoc();
	if(!DIL){
		return;
	}else{
		char pPath[200];
		string sFileName;
		if(DIL->getFilename().str().substr(0,5) == "/home"){
			sFileName = DIL->getFilename().str();
		}else{
			sFileName = DIL->getDirectory().str() + "/" + DIL->getFilename().str();
		}
		realpath(sFileName.c_str(), pPath);
		sFileName = string(sFileName);
		unsigned int numLine = DIL->getLine();
		errs() << "//---" << sFileName << ": " << numLine << "\n";
	}

	return;
}
/* Read source code from source file.
 */
string Support::ReadLine(const char *filename, int minline, int maxline){
    int lines,i=0;
    string temp;
    string contents = "";
    fstream file;
    file.open(filename,ios::in);
 
    if(file.fail())
    {
        return "Error 2: file do not exists!";
    }
    while(getline(file,temp))
    {
        i++;
        //errs() << temp << "\n";
        if((i >= minline - 1) && (i <= maxline+2 ) ){
            contents += temp + "\n";
        }
    }
    file.close();
    return contents;
}

/* Print source code of a loop.
 */
void Support::printLoopSrcCode(Loop *pLoop)
{
	int minlinenumber  = 0xfffffff;
	int maxlinenumber  = 0;
	string sFileName = "";
	for (Loop::block_iterator BI = pLoop->block_begin(); BI != pLoop->block_end(); BI++) {
		BasicBlock *BB = *BI;
		for (BasicBlock::iterator II = BB->begin(); II != BB->end(); II++) {
			Instruction *I = &*II;
			const DILocation * DIL = I->getDebugLoc();
			if(!DIL){
				continue;
			}else{
				char pPath[200];
				if(sFileName == ""){
					if(DIL->getFilename().str().substr(0,5) == "/home"){
						sFileName = DIL->getFilename().str();
					}else{
						sFileName = DIL->getDirectory().str() + "/" + DIL->getFilename().str();
					}
					realpath(sFileName.c_str(), pPath);
					sFileName = string(sFileName);
				}
				unsigned int numLine = DIL->getLine();
				//errs() << "//---" << sFileName << ": " << numLine << "\n";
				if(numLine < minlinenumber){
					minlinenumber = numLine;
				}
				if(numLine > maxlinenumber){
					maxlinenumber = numLine;
				}
			}
		}
	}
    errs()<< "--------------------------Source code:-----------------------------" << "\n";
	string lines = ReadLine( sFileName.c_str(), minlinenumber, maxlinenumber);
    errs()<< lines << "\n";
	errs()<< "-------------------------SRC Code end:-----------------------------" << "\n";
	return;
}
/* Print IR code of a loop.
 */
void Support::printLoopIRCode(Loop *pLoop)
{
    errs()<< "----------------------------IR code:-------------------------------" << "\n";
	pLoop->print(errs());
	errs()<< "--------------------------IR code end:-----------------------------" << "\n";
    for (Loop::block_iterator BI = pLoop->block_begin(); BI != pLoop->block_end(); BI++) {
		BasicBlock *BB = *BI;
        BB->print(errs());
    }
	errs()<< "--------------------------IR code end:-----------------------------" << "\n";
	return;
}
/* Get a function by name.
 */
Function * Support::SearchFunctionByName(Module & M, string strFunctionName)
{
	for(Module::iterator FI = M.begin(); FI != M.end(); FI ++ )
	{
		Function * f = &*FI;
		
		if(f->getName() == strFunctionName)
		{
			return f;
		}
	}
	return NULL;
}
/* Get a function id by name.
 */
int Support::GetFunctionID(Function *F) {

    if (F->begin() == F->end()) {
        return -1;
    }

    BasicBlock *EntryBB = &(F->getEntryBlock());

    if (EntryBB) {

        for (BasicBlock::iterator II = EntryBB->begin(); II != EntryBB->end(); II++) {
            Instruction *Inst = &*II;
            MDNode *Node = Inst->getMetadata("func_id");
            if (!Node) {
                continue;
            }
            assert(Node->getNumOperands() == 1);
            const Metadata *MD = Node->getOperand(0);
            if (auto *MDV = dyn_cast<ValueAsMetadata>(MD)) {
                Value *V = MDV->getValue();
                ConstantInt *CI = dyn_cast<ConstantInt>(V);
                assert(CI);
                return CI->getZExtValue();
            }
        }
    }

    return -1;
}
/* Get a function id by name.
 */
int Support::getPhiUser(set<Instruction *> *PhiUserSet, set<Instruction *> *ProcessedSet, Instruction * PhiIns){
	ProcessedSet->insert(PhiIns);
	//errs() << "Insert to ProcessedSet:";
	//PhiIns->print(errs());
	//errs() << "\n";
	for(auto U : PhiIns->users()){  // U is of type User*
		if (auto UI = dyn_cast<Instruction>(U)){
			if(ProcessedSet->find(UI) == ProcessedSet->end()){
				if(isa<PHINode>(UI)){
					getPhiUser(PhiUserSet, ProcessedSet, UI);
				}else{
					PhiUserSet->insert(UI);
					//errs() << "Insert to UserSet:";
					//UI->print(errs());
					//errs() << "\n";
				}
			}
		}
	}
	//errs() << ":" << UserSet->size() <<"\n";
	return PhiUserSet->size();
}
/* discarded function.
 */
bool Support::DefTracking(Instruction *UI, AllocaInst * AllocIns){
	//int i;
	UI->print(errs());
	errs()<<"\n";
	for (int i = 0; i < UI->getNumOperands(); i++) {

		auto Op = UI->getOperand(i);
		Op->print(errs());
		errs()<<"\n";
		if(Op && isa<AllocaInst>(Op)){
			AllocaInst *ai = dyn_cast<AllocaInst>(Op);
			if(ai == AllocIns){
				return true;
			}
		}else if(UI && isa<Instruction>(Op)){
			Instruction* OpIns = dyn_cast<Instruction>(Op);
			if(DefTracking(OpIns, AllocIns)){
				return true;
			}
		}
	}
	return false;
}
/* discarded function.
 */
bool Support::DependentOnItself(Instruction *DI){
	auto DesOp = DI->getOperand(1);
	auto SrcOp = DI->getOperand(0);
	if(!isa<AllocaInst>(DesOp)){
		return false;
	}
	AllocaInst *AllocIns = dyn_cast<AllocaInst>(DesOp);
	if(isa<Instruction>(SrcOp)){
		Instruction* SrcIns = dyn_cast<Instruction>(SrcOp);
		return DefTracking(SrcIns,AllocIns);
	}
	return false;
}
/* To determnie whether a instruction is in a loop's constant subloop.
 */
bool Support::IsInsideConstantSubLoop(Loop * pLoop, set<loopinfo>* SkipLoopSet, Instruction *UI, Instruction *DI){
	bool isInsideSubloop = false;
	/* Get all subloop.
	 */
	std::vector<Loop *> &subloopVect = pLoop->getSubLoopsVector();
	if (subloopVect.size() != 0) {
		for (std::vector<Loop *>::const_iterator SI = subloopVect.begin(); SI != subloopVect.end(); SI++) {
			(*SI)->print(errs());
			if (*SI != NULL) {
				loopinfo loop;
				string loop_label = (*SI)->getHeader()->getName();
				int num_loop = 0;
				sscanf(loop_label.c_str(),"loop_label:<%d>",&num_loop);
				loop.id = num_loop;
				loop.dep = (*SI)->getLoopDepth();
				/* If the subloop is a skipped constant loop.
				 */
				if (SkipLoopSet->find(loop) != SkipLoopSet->end()) {
					/* If the instruction is in the subloop.
				 	*/
					if((*SI)->contains(UI) && (*SI)->contains(DI)){
						errs() << "isInsideSubloop\n";
						isInsideSubloop=true;
					}
				}
			}
		}
	}
	return isInsideSubloop;
}
/* Get all user instruction of a store variable in loop.
 * @ Parameter: LifeInfoOfBackEdge contains all the live variable in the loop header.
 * @ Return: return val indicate how many cmp and switch have been found.
 * The instructions itself are return by the UserSet. 
 */
int Support::getUser(Loop *pLoop, Instruction *DI, set<Instruction *> *UserSet, DataLayout &DL, set<AllocaInst *>* LifeInfoOfBackEdge){
	bool loop_identified = false;
	#ifdef DEF_USE_DBG
	errs() << "A Store instruction:";
	DI->print(errs());
	errs() << "\n";
	#endif
	//LoadInst *LI = cast<LoadInst>(I);
	AllocaInst *allocins = nullptr;
	StoreInst * pStore = nullptr;
	bool isLiveInHeader = false;
	switch(DI->getOpcode()){
		/* We only care about the store instructions.
		 */
		case Instruction::Store:
			#ifdef DEF_USE_DBG
			errs() << "Store variable: ";
			DI->getOperand(1)->print(errs());
			errs() << "\n";
			#endif
			/* Handle cases that the variable exist in the struct.
			 */
			pStore = dyn_cast<StoreInst>(DI);
			if(Struct::isTargetInstruction(pStore, DL)){
				/* We get the struct it belongs to and the offset of its layout in the struct. 
			 	*/
				Struct::DecomposedGEP de;
				Struct::DecomposeGEPExpression(pStore->getPointerOperand(), de, DL);
#ifdef STRUCT_DBG 
				de.Base->print(errs());
				errs() << "\n" << "offset:" << de.StructOffset << "\n";
#endif
				/* We get all the cmp instructions and switch instructions which uses 
				 * the variable that the store instructon stores.
				 */
				if(isa<LoadInst>(de.Base)){
					const LoadInst *LoadInsForStore = dyn_cast<LoadInst>(de.Base);
#ifdef STRUCT_DBG 
					LoadInsForStore->getOperand(0)->print(errs());
#endif
					for (Loop::block_iterator BI = pLoop->block_begin(); BI != pLoop->block_end(); BI++) {
						BasicBlock *BB = *BI;
						for (BasicBlock::iterator II = BB->begin(); II != BB->end() && loop_identified == false; II++) {
							Instruction *UI = &*II;
							if(UI->getOpcode() == Instruction::Switch || UI->getOpcode() == Instruction::ICmp){
#ifdef STRUCT_DBG 
								errs() << "\n Cmp or Switch instruction(NumOperands:UI->" << UI->getNumOperands() << "):\n";
								UI->print(errs());
#endif
								for (int i = 0; i < UI->getNumOperands(); i++) {
									auto Op = UI->getOperand(i);
									/* There exists some cases that bit-extend a variable 
									 * for comparasion. We need to extract the original variable
									 * from it. 
									 */
									if(isa<ZExtInst>(Op)){
										ZExtInst * ZExtOp = dyn_cast<ZExtInst>(Op);
										Op = ZExtOp->getOperand(0);
									}
#ifdef STRUCT_DBG 
									errs() << "\n\tOperand("<< i <<"):\n";
									Op->print(errs());
#endif
									if(isa<LoadInst>(Op)){
										LoadInst *LoadOp = dyn_cast<LoadInst>(Op) ;
#ifdef STRUCT_DBG 
										errs() << "\nCmp use the load a instruction:\n";
										LoadOp->print(errs());
#endif
										/* For each variable that used for cmp of switch,
										 * get the we make sure that it is belong to a struct.
										 */
										if(Struct::isTargetInstruction(LoadOp, DL)){
#ifdef STRUCT_DBG 
											errs() << "\nisTargetInstruction:\n";
#endif
											/* Get the struct the var belongs to and the offset of its layout.
											 */
											Struct::DecomposedGEP us;
											Struct::DecomposeGEPExpression(LoadOp->getPointerOperand(), us, DL);
											if(isa<LoadInst>(us.Base)){	
												const LoadInst *LoadInsForCmp = dyn_cast<LoadInst>(us.Base);
#ifdef STRUCT_DBG 
												errs() << "\nLoadnInsForStore:\n";
												LoadInsForStore->getOperand(0)->print(errs());
												errs() << "\nLoadnInsForCmp:\n";
												LoadInsForCmp->getOperand(0)->print(errs());
												errs() << "\n";
#endif
												/* Add to the user set if the used variable and stored variable belong
												 * to the same struct and is of the same offset. 
												 */
												if((LoadInsForCmp->getOperand(0)->getType() == LoadInsForStore->getOperand(0)->getType()) 
														&& (us.StructOffset == de.StructOffset ) ){
													UserSet->insert(UI);
												}
											}
										}
									}
								}
							}
						}
					}
				}
				return UserSet->size();
			}
			/* Handle non-live variable
			 * It the variable not live in at the begin of the function header, 
			 * it is not supposed to be a statue variable, just skip it. 
			 */
			allocins = dyn_cast<AllocaInst>(DI->getOperand(1));
			if(LifeInfoOfBackEdge->find(allocins) != LifeInfoOfBackEdge->end()){
				errs() << "var not live in header\n";
				return 0;
			}
			/* Handle cases that the variable is just a simple variable.
			 */
			for(auto U : DI->getOperand(1)->users()){
				if(auto UI = dyn_cast<Instruction>(U)){
					if(pLoop->contains(UI)){	
						#ifdef DEF_USE_DBG
						errs() << "A use instruction:";
						UI->print(errs());
						errs() << "\n";
						#endif
						/* Get all the cmp and switch instructions which use the variable.
						 */
						switch(UI->getOpcode()){
							case Instruction::Load:
								for(auto UU : UI->users()){		
									if(auto UUI = dyn_cast<Instruction>(UU)){
										if(pLoop->contains(UUI)){
											if(UUI->getOpcode() == Instruction::Switch || UUI->getOpcode() == Instruction::ICmp){
												#ifdef DEF_USE_DBG
												errs() << "!Switch or ICmp instruction";
												UUI->print(errs());
												errs() << "\n";
												#endif
												UserSet->insert(UUI);
											}
										}
									}
									
								}
								break;
							case Instruction::Store:
								// Eliminate the cases that a variable flows to itself: var = f(var); 
#ifdef OPEN_DEPENDENT_ON_ITSELF_ELIMINATION	
								if(DependentOnItself(UI)){
									errs() << "Dependent on Itself\n";
									return 0;
								}
#endif
								break;
							default:
								#ifdef DEF_USE_DBG
								errs() << "!No handled instruction";
								errs() << "\n";
								#endif	
								break;
						}
					}
				}
			}
			break;
		case Instruction::Load:
			
			for(auto U : DI->users()){  // U is of type User*
				if (auto UI = dyn_cast<Instruction>(U)){
					#ifdef DEF_USE_DBG
					errs() << "A use instruction:";
					UI->print(errs());
					errs() << "\n";
					#endif

				}
			}
			break;
	}
	return UserSet->size();
}
/* Get all inevitable basic blocks in a loop.
 * A inevitable Basic Block is a basic basic that dominate by the back edges.
 * What is back edges: back edge is the edge to the loop header in the loop.
 */
set<BasicBlock *>* Support::getInevitableBBinLoop(Loop *pLoop, PostDominatorTree &pdt){
	BasicBlock *HB = &*pLoop->getHeader();
	set<BasicBlock *> BackEgeSet;
	BackEgeSet.clear();
	set<BasicBlock *>* InevitableBBSet = new set<BasicBlock *>;
	InevitableBBSet->clear();
	/* Get all the back edges in a loop.
	 */
	for (auto it = pred_begin(HB), et = pred_end(HB); it != et; ++it)
	{
		BasicBlock* predecessor = *it;
		if(pLoop->contains(predecessor)){
			BackEgeSet.insert(predecessor);
#ifdef INEV_BB
			errs() << "back eges :\n" ;
			predecessor->print(errs());
			errs() << "\n" ;
#endif
		}
	}
	/* Get the inevitable basic blocks.
	 */
	for (Loop::block_iterator BI = pLoop->block_begin(); BI != pLoop->block_end(); BI++) {
		BasicBlock *BB = *BI;
#ifdef INEV_BB
		errs() << "BasicBlock\n" ;
		BB->print(errs());
		errs() << "\n" ;
#endif
		bool isInevitableBB = true;
		for(auto BackEge : BackEgeSet){
			/* Postdominated by all backedge
			 */
			if(!(pdt.dominates(BackEge,BB))){
				isInevitableBB = false;
			}
		}
		if(isInevitableBB){
			InevitableBBSet->insert(BB);
		}
	}
	return InevitableBBSet;
}
