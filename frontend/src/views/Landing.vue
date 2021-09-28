<template>
  <board :boardId="boardId" v-if="boardId" />
</template>
<script lang="ts">
import { postBoard } from "@/api/board";
import { BoardDto, CreateBoardDto } from "@/types/board";
import { defineComponent } from "vue";
import { useQueryClient } from "vue-query";
import { useRoute, useRouter } from "vue-router";
import Board from "../components/Board.vue";

export default defineComponent({
  name: "Landing",
  components: {
    Board,
  },
  setup() {
    const router = useRouter();
    const route = useRoute();

    const createBoard = postBoard();
    const queryClient = useQueryClient();

    const boardId = route.params.boardId;

    if (!route.params.boardId) {
      console.log("foo");
      createBoard.mutate(
        {
          name: "New Board",
          goals: [],
          contents: [],
          phases: [],
        } as CreateBoardDto,
        {
          onSuccess: (data: BoardDto) => {
            queryClient.setQueryData(["board"], data);
            router.replace({
              name: "Landing",
              params: {
                boardId: data ? data.id : "",
              },
            });
          },
        }
      );
    }

    return { boardId };
  },
});
</script>
<style></style>
