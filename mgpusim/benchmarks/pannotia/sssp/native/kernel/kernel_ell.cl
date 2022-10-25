/************************************************************************************\
 *                                                                                  *
 * Copyright � 2014 Advanced Micro Devices, Inc.                                    *
 * All rights reserved.                                                             *
 *                                                                                  *
 * Redistribution and use in source and binary forms, with or without               *
 * modification, are permitted provided that the following are met:                 *
 *                                                                                  *
 * You must reproduce the above copyright notice.                                   *
 *                                                                                  *
 * Neither the name of the copyright holder nor the names of its contributors       *   
 * may be used to endorse or promote products derived from this software            *
 * without specific, prior, written permission from at least the copyright holder.  *
 *                                                                                  *
 * You must include the following terms in your license and/or other materials      *
 * provided with the software.                                                      * 
 *                                                                                  *  
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"      *
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE        *
 * IMPLIED WARRANTIES OF MERCHANTABILITY, NON-INFRINGEMENT, AND FITNESS FOR A       *
 * PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER        *
 * OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,         *
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT  * 
 * OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS      *
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN          *
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING  *
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY   *
 * OF SUCH DAMAGE.                                                                  *
 *                                                                                  *
 * Without limiting the foregoing, the software may implement third party           *  
 * technologies for which you must obtain licenses from parties other than AMD.     *  
 * You agree that AMD has not obtained or conveyed to you, and that you shall       *
 * be responsible for obtaining the rights to use and/or distribute the applicable  * 
 * underlying intellectual property rights related to the third party technologies. *  
 * These third party technologies are not licensed hereunder.                       *
 *                                                                                  *
 * If you use the software (in whole or in part), you shall adhere to all           *        
 * applicable U.S., European, and other export laws, including but not limited to   *
 * the U.S. Export Administration Regulations ("EAR") (15 C.F.R Sections 730-774),  *
 * and E.U. Council Regulation (EC) No 428/2009 of 5 May 2009.  Further, pursuant   *
 * to Section 740.6 of the EAR, you hereby certify that, except pursuant to a       *
 * license granted by the United States Department of Commerce Bureau of Industry   *
 * and Security or as otherwise permitted pursuant to a License Exception under     *
 * the U.S. Export Administration Regulations ("EAR"), you will not (1) export,     *
 * re-export or release to a national of a country in Country Groups D:1, E:1 or    * 
 * E:2 any restricted technology, software, or source code you receive hereunder,   * 
 * or (2) export to Country Groups D:1, E:1 or E:2 the direct product of such       *
 * technology or software, if such foreign produced direct product is subject to    * 
 * national security controls as identified on the Commerce Control List (currently * 
 * found in Supplement 1 to Part 774 of EAR).  For the most current Country Group   * 
 * listings, or for additional information about the EAR or your obligations under  *
 * those regulations, please refer to the U.S. Bureau of Industry and Security's    *
 * website at http://www.bis.doc.gov/.                                              *
 *                                                                                  *
\************************************************************************************/

#define BIG_NUM 99999999

/**
 * @brief   min.+
 * @param   num_nodes  number of vertices
 * @param   height     the height of the adjacency matrix (col-major)
 * @param   col        the col array
 * @param   data       the data array
 * @param   x          the input vector
 * @param   y          the output vector
 */
__kernel void
ell_min_dot_plus_kernel(   const  int    num_nodes,
                           const  int    height,
                         __global int   *col, 
                         __global int   *data, 
                         __global int   *x, 
                         __global int   *y)
{
    //get workitem id
    int  tid = get_global_id(0);
	
    if(tid < num_nodes){
	
        int mat_offset = tid;
        int min = x[tid]; 
        
		//the vertices process a row of matrix (col-major)
		for(int i = 0; i < height; i++){
            int mat_elem = data[mat_offset];
            int vec_elem = x[col[mat_offset]];
            if (mat_elem + vec_elem < min)
               min = mat_elem + vec_elem; 
            mat_offset += num_nodes;
        }
        y[tid]= min;
	}
}

/**
 * @brief   vector_init
 * @param   vector1    vector1
 * @param   vector2    vector2
 * @param   i          the source vertex id
 * @param   num_nodes  number of vertices
 */
__kernel void
vector_init(__global int *vector1,
            __global int *vector2,
               const int i, 
               const int num_nodes){

    //get my workitem id
	int tid = get_global_id(0);
	
    if (tid < num_nodes) {
        //if it is the source vertex
        if (tid == i){
            vector1[tid] = 0;
            vector2[tid] = 0;
        }
        //if it is a non-source vertex 
		else{
           vector1[tid] = BIG_NUM;
           vector2[tid] = BIG_NUM;
        }
	}

}

/**
 * @brief   vector_assign
 * @param   vector1    vector1
 * @param   vector2    vector2
 * @param   num_nodes  number of vertices
 */
__kernel void
vector_assign(__global int *vector1,
              __global int *vector2,
                const  int num_nodes){

    int tid = get_global_id(0);
    if (tid < num_nodes)
        vector1[tid] = vector2[tid];
}


/**
 * @brief   vector_diff
 * @param   vector1    vector1
 * @param   vector2    vector2
 * @param   stop       termination variable
 * @param   num_nodes  number of vertices
 */
__kernel void
vector_diff(__global int   *vector1,
            __global int   *vector2,
            __global bool  *stop,
	                 int    num_nodes){

	int tid = get_global_id(0);

    if (tid < num_nodes){
        if (vector2[tid] != vector1[tid])
            *stop = 1;
    }
}

