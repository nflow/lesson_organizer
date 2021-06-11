<template>
  <div class="w-full lg:w-5/6">
    <div class="bg-white shadow-md rounded my-6">
      <table class="min-w-max w-full table-auto">
        <tbody class="text-gray-600 text-sm font-light">
          <tr class="border-b border-gray-200 hover:bg-gray-100">
            <td class="py-3 px-6 text-left">
              <div class="flex items-center">
                <div class="mr-2">
                  <img
                    class="w-6 h-6 rounded-full"
                    src="https://randomuser.me/api/portraits/men/1.jpg"
                  />
                </div>
                <span>Eshal Rosas</span>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Options } from "vue-class-component";

@Options({
  props: ["heading"],
})
export default class List extends Vue {}
</script>
