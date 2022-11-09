<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="date" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="date" placeholder="结束时间"></el-date-picker>
      </el-form-item>
           <el-form-item label="学院" prop="college">
            <el-select v-model="searchInfo.college" clearable placeholder="请选择" @clear="()=>{searchInfo.college=undefined}">
              <el-option v-for="(item,key) in collegeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="姓名">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="学历" prop="qualifications">
            <el-select v-model="searchInfo.qualifications" clearable placeholder="请选择" @clear="()=>{searchInfo.qualifications=undefined}">
              <el-option v-for="(item,key) in education backgroundOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="岗位" prop="jobs">
            <el-select v-model="searchInfo.jobs" clearable placeholder="请选择" @clear="()=>{searchInfo.jobs=undefined}">
              <el-option v-for="(item,key) in jobsOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="学院" prop="college" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.college,collegeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="部门" prop="department" width="120" />
        <el-table-column align="left" label="姓名" prop="name" width="120" />
        <el-table-column align="left" label="性别" prop="gender" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.gender,genderOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="身份证号" prop="cardNumber" width="120" />
        <el-table-column align="left" label="职务" prop="position" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.position,positionOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="学历" prop="qualifications" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.qualifications,education backgroundOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="职称" prop="professional" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.professional,professional titleOptions) }}
            </template>
        </el-table-column>
         <el-table-column align="left" label="来校时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.hiredate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="岗位" prop="jobs" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.jobs,jobsOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateOaTeachersFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="学院:"  prop="college" >
          <el-select v-model="formData.college" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in collegeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="部门:"  prop="department" >
          <el-input v-model="formData.department" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:"  prop="name" >
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:"  prop="gender" >
          <el-select v-model="formData.gender" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="身份证号:"  prop="cardNumber" >
          <el-input v-model="formData.cardNumber" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="职务:"  prop="position" >
          <el-select v-model="formData.position" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in positionOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="学历:"  prop="qualifications" >
          <el-select v-model="formData.qualifications" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in education backgroundOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="职称:"  prop="professional" >
          <el-select v-model="formData.professional" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in professional titleOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="来校时间:"  prop="hiredate" >
          <el-date-picker v-model="formData.hiredate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="岗位:"  prop="jobs" >
          <el-select v-model="formData.jobs" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in jobsOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'OaTeachers'
}
</script>

<script setup>
import {
  createOaTeachers,
  deleteOaTeachers,
  deleteOaTeachersByIds,
  updateOaTeachers,
  findOaTeachers,
  getOaTeachersList
} from '@/api/oaTeachers'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const collegeOptions = ref([])
const genderOptions = ref([])
const positionOptions = ref([])
const education backgroundOptions = ref([])
const professional titleOptions = ref([])
const jobsOptions = ref([])
const formData = ref({
        college: undefined,
        department: '',
        name: '',
        gender: undefined,
        cardNumber: '',
        position: undefined,
        qualifications: undefined,
        professional: undefined,
        hiredate: new Date(),
        jobs: undefined,
        })

// 验证规则
const rule = reactive({
               college : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cardNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               position : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               qualifications : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               professional : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               hiredate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               jobs : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getOaTeachersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    collegeOptions.value = await getDictFunc('college')
    genderOptions.value = await getDictFunc('gender')
    positionOptions.value = await getDictFunc('position')
    education backgroundOptions.value = await getDictFunc('education background')
    professional titleOptions.value = await getDictFunc('professional title')
    jobsOptions.value = await getDictFunc('jobs')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteOaTeachersFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteOaTeachersByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateOaTeachersFunc = async(row) => {
    const res = await findOaTeachers({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reoaTeachers
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOaTeachersFunc = async (row) => {
    const res = await deleteOaTeachers({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        college: undefined,
        department: '',
        name: '',
        gender: undefined,
        cardNumber: '',
        position: undefined,
        qualifications: undefined,
        professional: undefined,
        hiredate: new Date(),
        jobs: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createOaTeachers(formData.value)
                  break
                case 'update':
                  res = await updateOaTeachers(formData.value)
                  break
                default:
                  res = await createOaTeachers(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
