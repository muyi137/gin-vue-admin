<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="身份证号:" prop="card_number">
          <el-input v-model="formData.card_number" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课头(几门课):" prop="classTime">
          <el-input v-model.number="formData.classTime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上报日期:" prop="reportTime">
          <el-date-picker v-model="formData.reportTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="k体:" prop="score">
          <el-input-number v-model="formData.score" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="k1体:" prop="score2">
          <el-input-number v-model="formData.score2" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="k1理:" prop="score1">
          <el-input-number v-model="formData.score1" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="k1机:" prop="score3">
          <el-input-number v-model="formData.score3" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="k2理:" prop="score4">
          <el-input-number v-model="formData.score4" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="k2机:" prop="score5">
          <el-input-number v-model="formData.score5" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="k3理:" prop="score6">
          <el-input-number v-model="formData.score6" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="k3机:" prop="score7">
          <el-input-number v-model="formData.score7" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="竞赛折算:" prop="score8">
          <el-input-number v-model="formData.score8" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="自定义1:" prop="score9">
          <el-input-number v-model="formData.score9" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="自定义2:" prop="score10">
          <el-input-number v-model="formData.score10" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="自定义3:" prop="score11">
          <el-input-number v-model="formData.score11" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="自定义4:" prop="score12">
          <el-input-number v-model="formData.score12" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="状态:" prop="status">
        <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['已审核','待审核']" :key="item" :label="item" :value="item" />
        </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  updateOaClassHour,
  findOaClassHour
} from '@/api/oaClassHour'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOaClassHour({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reoaClassHour
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
