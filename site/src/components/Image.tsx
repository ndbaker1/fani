import {
  Box,
  Button,
  CloseButton,
  DrawerActionTrigger,
  DrawerBackdrop,
  DrawerBody,
  DrawerCloseTrigger,
  DrawerContent,
  DrawerFooter,
  DrawerHeader,
  DrawerRoot,
  DrawerTitle,
  DrawerTrigger,
  Image,
  Stack,
} from "@chakra-ui/react";
import { FileObject } from "../gen/fani/v1/fani_pb";
import { Checkbox } from "./ui/checkbox";

type ImageProps = {
  data?: FileObject,
  source: string,
}

export default function(props: ImageProps) {
  return (
    <DrawerRoot size="full" contained={false}>
      <DrawerBackdrop />
      <Box position="relative">
        <DrawerTrigger asChild>
          <Image
            src={props.source}
            borderRadius={5}
            height={150}
          />
        </DrawerTrigger>
        <Checkbox
          top={0} left={0}
          position="absolute" />
      </Box>
      <DrawerContent position={"fixed"} top={0} left={0}>
        <DrawerHeader position="relative">
          <DrawerTitle>Drawer Title</DrawerTitle>
          <DrawerActionTrigger position="absolute" right={0} top={0} asChild>
            <CloseButton />
          </DrawerActionTrigger >
        </DrawerHeader>
        <DrawerBody>
          <Stack>
            <img src={props.source} />
            {JSON.stringify(props.data?.attributes)}
          </Stack>
        </DrawerBody>
        <DrawerFooter>
        </DrawerFooter >
        <DrawerCloseTrigger />
      </DrawerContent >
    </DrawerRoot >
  )
}

