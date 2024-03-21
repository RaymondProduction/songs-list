#include <gtk/gtk.h>

void init_combo_boxes(GtkBuilder *builder)
{

    GtkComboBoxText *comboboxtext1 = GTK_COMBO_BOX_TEXT(gtk_builder_get_object(builder, "select-song-1"));
    GtkComboBoxText *comboboxtext2 = GTK_COMBO_BOX_TEXT(gtk_builder_get_object(builder, "select-song-2"));
    GtkComboBoxText *comboboxtext3 = GTK_COMBO_BOX_TEXT(gtk_builder_get_object(builder, "select-song-3"));

    gtk_combo_box_text_append_text(comboboxtext1, "Option 1.1");
    gtk_combo_box_text_append_text(comboboxtext1, "Option 1.2");
    gtk_combo_box_text_append_text(comboboxtext1, "Option 1.3");

    gtk_combo_box_text_append_text(comboboxtext2, "Option 2.1");
    gtk_combo_box_text_append_text(comboboxtext2, "Option 2.2");
    gtk_combo_box_text_append_text(comboboxtext2, "Option 2.3");

    gtk_combo_box_text_append_text(comboboxtext3, "Option 3.1");
    gtk_combo_box_text_append_text(comboboxtext3, "Option 3.2");
    gtk_combo_box_text_append_text(comboboxtext3, "Option 3.3");
}

int main(int argc, char *argv[])
{
    gtk_init(&argc, &argv);

    GtkBuilder *builder = gtk_builder_new();
    if (gtk_builder_add_from_file(builder, "main.glade", NULL) == 0)
    {
        g_error("Error loading .glade file\n");
        return 1;
    }

    GtkWidget *window = GTK_WIDGET(gtk_builder_get_object(builder, "main-window"));
    g_signal_connect(window, "destroy", G_CALLBACK(gtk_main_quit), NULL);

    init_combo_boxes(builder);
    gtk_widget_show_all(window);

    gtk_main();
    return 0;
}