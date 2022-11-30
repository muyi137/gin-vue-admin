<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="身份证号">
         <el-input v-model="searchInfo.card_number" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="上报日期">
            
            <el-date-picker v-model="searchInfo.startReportTime" type="datetime" placeholder="搜索条件（起）"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endReportTime" type="datetime" placeholder="搜索条件（止）"></el-date-picker>

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
            <el-upload
          class="excel-btn"
          :action="`${path}/oaClassHour/importOaClassHour`"
          :headers="{'x-token':userStore.token}"
          :on-success="getTableData"
          :show-file-list="false"
        >
        <el-button size="small" type="primary" icon="upload" style="margin-left: 10px;">导入</el-button>
        </el-upload>
        <el-button class="excel-btn" size="small" type="primary" icon="download" @click="exportExc('课时')" style="margin-left: 10px;">导出</el-button>
        <el-button class="excel-btn" size="small" type="success" icon="download" @click="downloadExcelTemplate()">下载模板</el-button>
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
         
        <el-table-column align="left" label="操作" width="160">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateOaClassHourFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        <el-table-column align="left" label="身份证号" prop="card_number" width="120" />
        <el-table-column align="left" label="课头(几门课)" prop="classTime" width="120" />
         <el-table-column align="left" label="上报日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.reportTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="k体" prop="score" width="120" />
        <el-table-column align="left" label="k1体" prop="score2" width="120" />
        <el-table-column align="left" label="k1理" prop="score1" width="120" />
        <el-table-column align="left" label="k1机" prop="score3" width="120" />
        <el-table-column align="left" label="k2理" prop="score4" width="120" />
        <el-table-column align="left" label="k2机" prop="score5" width="120" />
        <el-table-column align="left" label="k3理" prop="score6" width="120" />
        <el-table-column align="left" label="k3机" prop="score7" width="120" />
        <el-table-column align="left" label="竞赛折算" prop="score8" width="120" />
        <el-table-column align="left" label="自定义1" prop="score9" width="120" />
        <el-table-column align="left" label="自定义2" prop="score10" width="120" />
        <el-table-column align="left" label="自定义3" prop="score11" width="120" />
        <el-table-column align="left" label="自定义4" prop="score12" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
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
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item label="身份证号:"  prop="card_number" style="width:40%">
          <el-input v-model="formData.card_number" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课头(几门课):"  prop="classTime" style="width:20%"  >
          <el-input v-model.number="formData.classTime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上报日期:"  prop="reportTime" >
          <el-date-picker v-model="formData.reportTime" type="date" style="width:40%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="k体:"  prop="score" >
          <el-input-number v-model="formData.score"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="k1体:"  prop="score2" >
          <el-input-number v-model="formData.score2"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="k1理:"  prop="score1" >
          <el-input-number v-model="formData.score1"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="k1机:"  prop="score3" >
          <el-input-number v-model="formData.score3"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="k2理:"  prop="score4" >
          <el-input-number v-model="formData.score4"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="k2机:"  prop="score5" >
          <el-input-number v-model="formData.score5"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="k3理:"  prop="score6" >
          <el-input-number v-model="formData.score6"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="k3机:"  prop="score7" >
          <el-input-number v-model="formData.score7"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="竞赛折算:"  prop="score8" >
          <el-input-number v-model="formData.score8"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="自定义1:"  prop="score9" >
          <el-input-number v-model="formData.score9"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="自定义2:"  prop="score10" >
          <el-input-number v-model="formData.score10"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="自定义3:"  prop="score11" >
          <el-input-number v-model="formData.score11"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="自定义4:"  prop="score12" >
          <el-input-number v-model="formData.score12"  style="width:20%" :precision="2" :clearable="true"  />
        </el-form-item>
        <!-- <el-form-item label="状态:"  prop="status" >
            <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true" >
               <el-option v-for="item in ['已审核','待审核']" :key="item" :label="item" :value="item" />
            </el-select>
        </el-form-item> -->
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
  name: 'OaClassHour'
}
</script>

<script setup>
import {
  createOaClassHour,
  deleteOaClassHour,
  deleteOaClassHourByIds,
  updateOaClassHour,
  findOaClassHour,
  getOaClassHourList
} from '@/api/oaClassHour'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
import * as XLSX from 'xlsx'

const path = ref(import.meta.env.VITE_BASE_API)
const userStore = useUserStore()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        card_number: '',
        classTime: 0,
        reportTime: new Date(),
        score: 0,
        score2: 0,
        score1: 0,
        score3: 0,
        score4: 0,
        score5: 0,
        score6: 0,
        score7: 0,
        score8: 0,
        score9: 0,
        score10: 0,
        score11: 0,
        score12: 0,
        status : '待审核',
        })

// 验证规则
const rule = reactive({
               card_number : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               classTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               reportTime : [{
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
  const table = await getOaClassHourList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteOaClassHourFunc(row)
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
      const res = await deleteOaClassHourByIds({ ids })
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
const updateOaClassHourFunc = async(row) => {
    const res = await findOaClassHour({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reoaClassHour
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOaClassHourFunc = async (row) => {
    const res = await deleteOaClassHour({ ID: row.ID })
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

const exportExc = async(name) => {
  const table = await getOaSalaryList({ page: page.value, pageSize: 9999999, ...searchInfo.value })
  if (table.code === 0) {
    var data = [
      ['身份证号', '课头(几门课)', '上报日期', 'k体', 'k1体', 'k1理', 'k1机', 'k2理', 'k2机', 'k3理', 'k3机', '竞赛折算', '自定义1', '自定义2', '自定义3', '自定义4']
    ]

    var res = table.data.list
    // const header = ['身份证号', '月份', '绩效值']
 
    for (var i = 0; i < res.length; i++) {
      var params = [
        res[i].card_number,
        res[i].classTime,
        res[i].reportTime,
        res[i].score,
        res[i].score2,
        res[i].score3,
        res[i].score4,
        res[i].score5,
        res[i].score6,
        res[i].score7,
        res[i].score8,
        res[i].score9,
        res[i].score10,
        res[i].score11,
        res[i].score12
      ]
      data[i + 1] = params
    }
    // console.log(JSON.stringify(data))
    var wb = XLSX.utils.aoa_to_sheet(data, {
      // header: header,
      raw: true
    }
    )
    const newWorkBook = XLSX.utils.book_new()
    XLSX.utils.book_append_sheet(newWorkBook, wb, 'SheetJS')
    var myDate = new Date()
    var wbout = XLSX.writeFile(newWorkBook, name + myDate.toLocaleString() + '.xlsx')

    // var wbout = XLSX.write(wb, { bookType: 'xlsx', bookSST: true, type: 'array' }) 
    // try {
    //   FileSaver.saveAs(new Blob([wbout], { type: 'application/octet-stream' }), name + myDate.toLocaleString() + '.xlsx')
    // } catch (e) {
    //   if (typeof console !== 'undefined') console.log(e, wbout) 
    // }

    return wbout
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
        card_number: '',
        classTime: 0,
        reportTime: new Date(),
        score: 0,
        score2: 0,
        score1: 0,
        score3: 0,
        score4: 0,
        score5: 0,
        score6: 0,
        score7: 0,
        score8: 0,
        score9: 0,
        score10: 0,
        score11: 0,
        score12: 0,
        status : '待审核',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createOaClassHour(formData.value)
                  break
                case 'update':
                  res = await updateOaClassHour(formData.value)
                  break
                default:
                  res = await createOaClassHour(formData.value)
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
