import {AfterViewInit, Directive, ElementRef, Input} from '@angular/core';

@Directive({
    selector: '[appImage]'
})
export class ImageDirective implements AfterViewInit {

    @Input() src;
    @Input() type;

    constructor(private imageRef: ElementRef) {
    }

    ngAfterViewInit(): void {
        const img = new Image();
        img.onload = () => {
            this.setImage('./assets/media/images/' + this.src);
        };

        img.onerror = () => {
            this.setImage('./assets/media/plasma/default.jpg');
        };

        img.src = './assets/media/images/' + this.src;
    }

    private setImage(src: string) {
        this.imageRef.nativeElement.setAttribute('src', src);
    }
}
