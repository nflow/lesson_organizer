<template>
  <div class="tw-flex tw-flex-col tw-bg-gray-300 tw-h-full tw-w-full tw-flex">
    <div class="tw-flex-initial tw-flex tw-flex-col tw-bg-gray-200 tw-m-0">
      <div
        class="
          tw-flex-initial
          tw-bg-gradient-to-br
          tw-from-blue
          tw-to-turquoise
          tw-p-4
          tw-shadow-mds
          tw-flex
          sm:tw-flex-row
          tw-flex-col tw-gap-2
        "
      >
        <div
          class="
            tw-flex-initial
            tw-items-center
            tw-p-2
            tw-rounded
            tw-text-white
            tw-shadow-xl
            tw-bg-gradient-to-br
            tw-from-red-300
            tw-to-orange
            tw-text-xl
            tw-text-center
            tw-font-logo
          "
        >
          Lesson Organizer
        </div>
      </div>
    </div>
    <div class="tw-m-auto tw-w-1/4">
      <form @submit.prevent.stop="onCreateBoard">
        <q-card class="sm:tw-w-full md:tw-w-10/12">
          <q-card-section>
            <span class="tw-font-bold tw-text-xl">Create Board</span>
          </q-card-section>
          <q-card-section>
            <q-input
              ref="boardNameRef"
              v-model="boardModel.name"
              square
              outlined
              lazy-rules
              :rules="boardNameRules"
              label="Board Name"
            />
          </q-card-section>
          <q-card-section class="tw-space-x-2">
            <q-btn
              color="primary"
              :loading="createBoard.isLoading.value"
              type="submit"
              label="Create Board"
            />
          </q-card-section>
        </q-card>
      </form>
    </div>
  </div>
</template>
<script lang="ts">
import { postBoard } from "@/api/board";
import { BoardDto, CreateBoardDto } from "@/types/board";
import { defineComponent, Ref, ref } from "vue";
import { useRouter } from "vue-router";
export default defineComponent({
  name: "Landing",
  setup() {
    const router = useRouter();

    const createBoard = postBoard();

    const boardModel: Ref<CreateBoardDto> = ref({ name: "" });

    const boardNameRef = ref<any>({});

    const onCreateBoard = () => {
      boardNameRef.value.validate();

      if (!boardNameRef.value.hasError) {
        createBoard.mutate(boardModel.value, {
          onSuccess: (data: BoardDto) => {
            router.replace({
              name: "Board",
              params: {
                boardId: data ? data.id : "",
              },
            });
          },
        });
      }
    };

    return {
      boardModel,
      createBoard,
      boardNameRef,
      boardNameRules: [
        (val: string | undefined) =>
          (val && val.length > 0) || "Please type something",
      ],
      onCreateBoard,
    };
  },
});
</script>
<style></style>
